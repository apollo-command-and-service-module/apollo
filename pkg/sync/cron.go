package sync

import (
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"github.com/apollo-command-and-service-module/apollo/pkg/job"
	"github.com/apollo-command-and-service-module/apollo/pkg/repo"
	"github.com/apollo-command-and-service-module/apollo/pkg/viper"
	"gopkg.in/robfig/cron.v2"
	"log"
	"time"
)

var scheduler cron.Cron

func StartScheduler(jobFrequency *string, jobQueue chan job.Job) {
	scheduler := cron.New()
	scheduler.Start()
	scheduler.AddFunc(*jobFrequency, func() { AddJob(jobQueue) })
}

func StopScheduler() {
	scheduler.Stop()
}

func AddJob(jobQueue chan job.Job) {

	// Read Config for Services
	services := viper.Services()
	log.Print("TODO: Sync lunar-module")

	for _, s := range services {
		liftoff := pkg.FormatDate(time.Now())
		id := pkg.IdGenerator()

		setting := repo.Repo{
			Url:        s.Url,
			Branch:     s.Branch,
			ConfigFile: s.Config,
		}

		job := job.Job{
			Id:          id,
			Name:        s.Name,
			Status:      job.StatusQueued,
			CurrentDate: liftoff,
			Repo:        setting,
		}

		// Add service to job queue
		jobQueue <- job
	}
}
