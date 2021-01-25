package main

import (
	"flag"
	"github.com/apollo-command-and-service-module/apollo/pkg/job"
	cron "github.com/apollo-command-and-service-module/apollo/pkg/sync"
	"log"
	"net/http"
)

func main() {
	var (
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
		maxQueueSize = flag.Int("max_queue_size", 10, "The size of job queue")
		jobFrequency = flag.String("job_frequency", "@every 0h1m0s", "The start job frequency")
		port         = flag.String("port", "5000", "The server port")
		source       = flag.String("source", "disk", "AGC Storage Mode (disk, s3, or env) ")
		filename     = flag.String("filename", "agc", "The Apollo Guidance Configuration(AGC) Yaml File")
		directory    = flag.String("directory", "./agc", "AGC Storage Directory")
		mainEngine   = flag.String("url", "https://github.com/apollo-command-and-service-module/orbit.git", "The Apollo Guidance Configuration Repository")
		engineBranch = flag.String("branch", "main", "The Apollo Guidance Configuration Repository Branch")
	)
	flag.Parse()

	// Create the job queue.
	jobQueue := make(chan job.Job, *maxQueueSize)

	// Start the worker dispatcher.
	dispatcher := job.NewDispatcher(jobQueue, *maxWorkers)
	dispatcher.Run()

	if *source != "" {
		// Start the Cron Scheduler
		log.Printf("StartMgs: Apollo, Houston. You're good at 1 minute.")
		setConfig := cron.SetConfig(string(*source), string(*directory), string(*mainEngine), string(*filename), string(*engineBranch))
		setConfig.StartScheduler(jobFrequency, jobQueue)
	}
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
