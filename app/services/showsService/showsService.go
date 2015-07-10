package showsService

import (
	"strconv"
	"time"

	tvdb "github.com/garfunkel/go-tvdb"
	"github.com/kahoona77/emerald/app/models"
	"github.com/kahoona77/emerald/app/services/mongo"
)

const SHOWS_REPO = "shows"
const EPISODES_REPO = "episodes"
const dateFormat = "2006-01-02"

func FindAllShows() ([]models.Show, error) {
	var shows []models.Show
	err := mongo.All(SHOWS_REPO, &shows)
	return shows, err
}

func SaveShow(show *models.Show) error {
	_, err := mongo.Save(SHOWS_REPO, show.Id, show)
	return err
}

func DeleteShow(show *models.Show) error {
	err := mongo.Remove(SHOWS_REPO, show.Id)
	return err
}

func SearchShow(query string) ([]models.Show, error) {
	var shows []models.Show

	results, err := tvdb.GetSeries(query)
	if err != nil {
		return nil, err
	}

	shows = make([]models.Show, len(results.Series), len(results.Series))

	for i := range results.Series {
		shows[i] = showFromSeries(results.Series[i])
	}

	return shows, nil
}

func showFromSeries(series *tvdb.Series) models.Show {
	show := models.Show{}
	show.Name = series.SeriesName
	show.SearchName = series.SeriesName
	show.Folder = series.SeriesName
	show.Id = strconv.Itoa(int(series.ID))
	show.Banner = series.Banner
	show.Poster = series.Poster
	show.FirstAired = parseDate(series.FirstAired)
	show.Overview = series.Overview

	return show
}

func parseDate(date string) time.Time {
	t, _ := time.Parse(dateFormat, date)
	return t
}

func UpdateShow(show *models.Show) (*models.Show, error) {
	seriesID, err := strconv.Atoi(show.Id)
	series, err := tvdb.GetSeriesByID(uint64(seriesID))
	if err != nil {
		return nil, err
	}

	//update fields & save show
	show.Banner = series.Banner
	show.Poster = series.Poster
	show.FirstAired = parseDate(series.FirstAired)
	show.Overview = series.Overview
	_, err = mongo.Save(SHOWS_REPO, show.Id, show)
	if err != nil {
		return nil, err
	}

	//update episodes
	updateEpisodes(show)

	return show, nil
}
