package infrastructure

import (
	domain "blogs/Domain"
	"context"
	"log"

	"github.com/robfig/cron/v3"
)


type BackgroundTask struct {
	usecase domain.SignupUseCase
}

func NewBackgroundTask(userusecase domain.SignupUseCase) *BackgroundTask {
	return &BackgroundTask{
		usecase: userusecase,
	}
}

func (bt *BackgroundTask) StartCronJob() {
    c := cron.New()

    // Add a cron job to run every minute (using the appropriate cron expression)
    _, err := c.AddFunc("@every 360h", func() {
        err := bt.usecase.DeleteOldUnverifiedUsers(context.Background() ,15) // Delete users older than 15 day (adjust this parameter if needed)
        if err != nil {
            log.Println(err)
        } else {
            log.Println("Unverified users deleted successfully")
        }
    })

    if err != nil {
        log.Println("Error adding cron job:", err)
    } else {
        // Start the cron scheduler
        c.Start()
    }
}