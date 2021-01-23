package sync

import (
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"github.com/apollo-command-and-service-module/apollo/pkg/job"
	"github.com/apollo-command-and-service-module/apollo/pkg/logging"
	"github.com/apollo-command-and-service-module/apollo/pkg/repo"
	"github.com/apollo-command-and-service-module/apollo/pkg/viper"
	"gopkg.in/robfig/cron.v2"
	"log"
	"os"
	"time"
)

var scheduler cron.Cron

type Config struct {
	Source        string
	Directory     string
	Engine        string
	Filename      string
	Branch        string
}

func SetConfig( source string, directory string, engine string, filename string, branch string) *Config {

	return &Config{
		Source:        source,
		Directory:     directory,
		Engine:        engine,
		Filename:      filename,
		Branch:        branch,
	}
}

func (c *Config) StartScheduler(jobFrequency *string, jobQueue chan job.Job) {

	switch c.Source {
	case "s3":
		log.Printf("TODO: AGC Configuration source s3")
	case "env":
		log.Printf("AGC Configuration source env")
		c.Directory = os.Getenv("AGC_PATH")
		c.Filename = os.Getenv("AGC_FILE")
	default:
		_, err := repo.CleanDiskStorage(c.Directory)
		logging.CheckIfError(err)

		repo.DiskStorage(c.Directory, c.Engine)
	}

	scheduler := cron.New()
	scheduler.Start()
	setAgc := viper.SetAgc(c.Filename,c.Directory)
	scheduler.AddFunc(*jobFrequency, func() { AddJob(jobQueue,setAgc, c.Source) })
}

func StopScheduler() {
	scheduler.Stop()
}

func AddJob(jobQueue chan job.Job, agc *viper.Agc, source string) {

	if source == "disk" {
		repo.Pull(agc.Directory)
	}
	services := agc.Services()

	// TODO: Sync lunar-module

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
