package showsService

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/kahoona77/emerald/models"

	"labix.org/v2/mgo/bson"
)

var (
	seRegex, _ = regexp.Compile(`(.*?)[.\s][sS](\d{2})[eE](\d{2}).*`)
	xRegex, _  = regexp.Compile(`(.*?)[.\s](\d{1,2})[xX](\d{2}).*`)
)

// ShowInfo Info about a show
type ShowInfo struct {
	Name    string
	Season  int
	Episode int
}

// ScanDownloadDir tries to move episodes from the download-folder
func (s *ShowsService) ScanDownloadDir(settings *models.XtvSettings) {
	// iterate over files in downlod-Dir
	err := filepath.Walk(settings.DownloadDir, func(file string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			s.MoveEpisode(file, settings, false)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error while updating episodes: %v", err)
	}

	UpdateKodi(settings)
}

// MoveEpisode moves the epsiode to its season folder
func (s *ShowsService) MoveEpisode(file string, settings *models.XtvSettings, updateKodi bool) {
	info := s.parseShow(file)

	if info != nil {
		show, episode := s.getShowData(info)

		if show == nil || episode == nil {
			//error
			return
		}

		// create output file
		fileEnding := file[strings.LastIndex(file, "."):]
		destinationFolder := settings.ShowsFolder + "/" + show.Folder + "/Season " + strconv.Itoa(int(episode.SeasonNumber)) + "/"
		fileName := s.sanitizeFilename(show.Name + " - " + strconv.Itoa(int(episode.SeasonNumber)) + "x" + fmt.Sprintf("%0.2d", episode.EpisodeNumber) + " - " + episode.Name)

		//create all directories
		err := os.MkdirAll(destinationFolder, 0755)
		if err != nil {
			log.Printf("Error making dir: %s", err)
		}

		//move epsiode to destination
		srcFile := filepath.FromSlash(file)
		destFile := filepath.FromSlash(destinationFolder + fileName + fileEnding)
		err = os.Rename(srcFile, destFile)
		if err != nil {
			log.Printf("Error while moving epsiode to destination: %s", err)
			return
		}

		log.Printf("Moved Episode '%s' to folder '%s'", fileName, destinationFolder)
	}

	if updateKodi == true {
		UpdateKodi(settings)
	}
}

func (s *ShowsService) getShowData(info *ShowInfo) (*models.Show, *models.Episode) {
	// find show
	var shows []models.Show
	query := bson.M{"searchName": info.Name}
	err := s.Mongo.FindWithQuery(SHOWS_REPO, &query, &shows)
	if err != nil || len(shows) <= 0 {
		log.Printf("could not find show: %v", info.Name)
		return nil, nil
	}
	show := shows[0]

	// find Episode
	var episode *models.Episode
	seasons, err := s.LoadEpisodes(show.Id)
	episodes := seasons[strconv.Itoa(int(info.Season))]
	for i := range episodes {
		if int(episodes[i].EpisodeNumber) == info.Episode {
			episode = &episodes[i]
		}
	}

	if episode == nil {
		log.Printf("could not find episode: %v for show %v", info.Episode, info.Name)
		return nil, nil
	}

	return &show, episode
}

func (s *ShowsService) sanitizeFilename(filename string) string {
	// Remove all strange characters
	seps, err := regexp.Compile(`[&_=+:]`)
	if err == nil {
		filename = seps.ReplaceAllString(filename, "")
	}

	return filename
}

func (s *ShowsService) parseShow(absoluteFile string) *ShowInfo {
	info := new(ShowInfo)

	// Replace all \ with /
	absoluteFile = strings.Replace(absoluteFile, "\\", "/", -1)

	// cut off the path
	file := absoluteFile[strings.LastIndex(absoluteFile, "/")+1:]

	// Replace all _ with dots
	file = strings.Replace(file, "_", ".", -1)

	result := seRegex.FindStringSubmatch(file)
	if result != nil {
		info.Name = strings.Replace(result[1], ".", " ", -1)
		info.Season, _ = strconv.Atoi(result[2])
		info.Episode, _ = strconv.Atoi(result[3])
	} else {
		// try othe rpattern
		result = xRegex.FindStringSubmatch(file)
		if result != nil {
			info.Name = strings.Replace(result[1], ".", " ", -1)
			info.Season, _ = strconv.Atoi(result[2])
			info.Episode, _ = strconv.Atoi(result[3])
		} else {
			return nil
		}
	}

	return info
}

// UpdateKodi sned a request to Kodi to update its database
func UpdateKodi(settings *models.XtvSettings) {
	//connect
	conn, err := net.Dial("tcp", settings.KodiAddress)
	if err != nil {
		log.Printf("Error while connecting to Kodi: %v", err)
		return
	}
	defer conn.Close()

	msg := `{"id":1,"method":"VideoLibrary.Scan","params":[],"jsonrpc":"2.0"}`
	// json.NewEncoder(conn).Encode(data)
	if _, err = conn.Write([]byte(msg)); err != nil {
		log.Printf("Error while sending update command to Kodi: %s", err)
	}
}
