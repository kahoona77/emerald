package jobs

import (
	"log"

	"github.com/kahoona77/emerald/services/showsService"
	"github.com/robfig/cron"
)

//UpdateJob updates all episodes in the night
type UpdateJob struct {
	ShowsService *showsService.ShowsService `inject:""`
}

//StartJob starts the UpdateJob
func (uj *UpdateJob) StartJob() {
	c := cron.New()
	c.AddFunc("@daily", uj.updateShows)
	c.Start()
}

func (uj *UpdateJob) updateShows() {
	shows, err := uj.ShowsService.FindAllShows()
	if err != nil {
		log.Printf("Error while updating shows: %v", err)
	}
	for _, show := range shows {
		uj.ShowsService.UpdateShow(&show)
	}
	log.Printf("Updated %d shows", len(shows))
}
