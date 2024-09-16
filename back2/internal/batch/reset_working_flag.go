package batch

import (
	"back2/internal/domain/repository"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func ScheduleResetWorkingFlag(castRepo repository.CastRepository) {
	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Tokyo", 9*60*60)))

	log.Println("Scheduling reset WorkingFlg job")

	_, err := c.AddFunc("0 6 * * *", func() {
		log.Println("Starting to reset Cast WorkingFlg to 0")
		err := castRepo.ResetAllWorkingFlags()
		if err != nil {
			log.Printf("Error resetting WorkingFlg: %v", err)
		} else {
			log.Println("Successfully reset all Cast WorkingFlg to 0")
		}
	})

	if err != nil {
		log.Fatalf("Error scheduling reset WorkingFlg job: %v", err)
	}

	log.Println("Reset WorkingFlg job scheduled successfully")

	c.Start()
	log.Println("Cron job started")
}
