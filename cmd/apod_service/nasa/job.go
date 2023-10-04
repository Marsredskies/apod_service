package nasa

import (
	"context"
	"log"
	"time"
)

// Kind of a cron job restarting every day or specified interval
// Checks if latest image exists in db before fetching
func (n *NasaClient) DoJobFetchAndSaveImages() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		latest, err := n.db.GetByDate(ctx, time.Now().Format(dateFormat))
		if err == nil && latest == nil {
			err = n.FetchAndSaveAPOD(ctx, time.Now())
			if err != nil {
				log.Printf("image fetch failed: %v ", err)
				continue
			}
		} else {
			if err != nil {
				log.Printf("failed to get latest image from db: %v", err)
			}
			if latest != nil {
				log.Printf("%s image already exists", latest.Date)
			}
		}

		time.Sleep(time.Duration(n.fetchIntervalHours) * time.Hour)
	}
}
