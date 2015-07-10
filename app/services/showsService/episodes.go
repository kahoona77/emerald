package showsService

import (
	"strconv"
	"time"

	"labix.org/v2/mgo/bson"

	tvdb "github.com/garfunkel/go-tvdb"
	"github.com/kahoona77/emerald/app/models"
	"github.com/kahoona77/emerald/app/services/mongo"
)

func LoadEpisodes(showId string) (map[string][]models.Episode, error) {
	queryObject := bson.M{"showId": showId}

	var episodes []models.Episode
	err := mongo.FindWithQuery(EPISODES_REPO, &queryObject, &episodes)
	if err != nil {
		return nil, err
	}

	seasons := make(map[string][]models.Episode)
	for _, episode := range episodes {
		seasonEps := seasons[strconv.Itoa(int(episode.SeasonNumber))]
		if seasonEps == nil {
			seasonEps = make([]models.Episode, 0)
		}

		seasonEps = append(seasonEps, episode)

		seasons[strconv.Itoa(int(episode.SeasonNumber))] = seasonEps
	}

	return seasons, nil
}

func updateEpisodes(show *models.Show) ([]*models.Episode, error) {
	//delete current episodes
	deleteEpisodes(show)

	seriesID, err := strconv.Atoi(show.Id)
	series, err := tvdb.GetSeriesByID(uint64(seriesID))
	if err != nil {
		return nil, err
	}

	err = series.GetDetail()
	if err != nil {
		return nil, err
	}

	episodes := make([]*models.Episode, 0)
	for _, eps := range series.Seasons {
		for i := range eps {
			episode := episodeFromSeriesEpisode(eps[i])
			err = saveEpisode(episode)
			if err != nil {
				return nil, err
			}
			episodes = append(episodes, episode)
		}
	}

	return episodes, nil
}

func deleteEpisodes(show *models.Show) (int, error) {
	removeQuery := bson.M{"showId": show.Id}
	info, err := mongo.RemoveAll(EPISODES_REPO, &removeQuery)
	return info.Removed, err
}

func saveEpisode(episode *models.Episode) error {
	_, err := mongo.Save(EPISODES_REPO, episode.Id, episode)
	return err
}

func episodeFromSeriesEpisode(seriesEpisode *tvdb.Episode) *models.Episode {
	episode := models.Episode{}
	episode.Id = strconv.Itoa(int(seriesEpisode.ID))
	episode.ShowId = strconv.Itoa(int(seriesEpisode.SeriesID))
	episode.Name = seriesEpisode.EpisodeName
	episode.FirstAired = parseDate(seriesEpisode.FirstAired)
	episode.Overview = seriesEpisode.Overview
	episode.Filename = seriesEpisode.Filename
	episode.EpisodeNumber = seriesEpisode.EpisodeNumber
	episode.SeasonNumber = seriesEpisode.SeasonNumber

	return &episode
}

func LoadRecentEpisodes(duration string) ([]models.RecentEpisode, error) {
	days, _ := strconv.Atoi(duration)
	dealy, _ := time.ParseDuration("-" + strconv.Itoa(days*24) + "h")
	pastDate := time.Now().Add(dealy)
	query := bson.M{"firstAired": bson.M{"$gt": pastDate, "$lt": time.Now()}}

	var episodes []models.Episode
	err := mongo.FindWithQuery(EPISODES_REPO, &query, &episodes)
	if err != nil {
		return nil, err
	}

	recentEpisodes := make([]models.RecentEpisode, 0)
	for _, episode := range episodes {
		var show models.Show
		mongo.FindById(SHOWS_REPO, episode.ShowId, &show)
		recentEpisode := models.RecentEpisode{}
		recentEpisode.Episode = episode
		recentEpisode.Show = show
		recentEpisodes = append(recentEpisodes, recentEpisode)
	}

	return recentEpisodes, nil
}
