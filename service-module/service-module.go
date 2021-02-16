package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/apollo-command-and-service-module/apollo/pkg/config"
	"github.com/apollo-command-and-service-module/apollo/pkg/job"
	"github.com/apollo-command-and-service-module/apollo/pkg/logging"
	cron "github.com/apollo-command-and-service-module/apollo/pkg/sync"
	"github.com/apollo-command-and-service-module/apollo/service-module/api/healthcheck"
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
		log          *logging.Logger
		textConfig   config.Configurations = config.NewTextConfig()
	)
	flag.Parse()

	//TODO: decide whether to use a 3rd party mux later.
	//for now using the one in the standard library
	*port = ":" + *port
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthcheck.Index)

	//define a server struct
	srv := &http.Server{
		Handler:      mux,
		Addr:         *port,
		WriteTimeout: 50 * time.Second,
		ReadTimeout:  50 * time.Second,
	}

	go func() {
		log = logging.NewConsole(true)
		textFormat := textConfig.SetTextFormatting("cyan")
		// Create the job queue.
		jobQueue := make(chan job.Job, *maxQueueSize)
		// Start the worker dispatcher.
		dispatcher := job.NewDispatcher(jobQueue, *maxWorkers)
		dispatcher.Run()

		if *source != "" {
			// Start the Cron Scheduler
			fmt.Println(textFormat.Color, "StartMgs: Apollo, Houston. You're good at 1 minute.", textFormat.Reset)
			setConfig := cron.SetConfig(string(*source), string(*directory), string(*mainEngine), string(*filename), string(*engineBranch))
			setConfig.StartScheduler(jobFrequency, jobQueue)
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Printf("Received Terminate, graceful shutdown %s", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(tc)
}
