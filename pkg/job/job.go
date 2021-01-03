package job

import (
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"github.com/apollo-command-and-service-module/apollo/pkg/repo"
	"net/http"
	"time"
)

type StatusString string

const (
	StatusQueued    StatusString = "queued"
	StatusRunning   StatusString = "running"
	StatusFailed    StatusString = "failed"
	StatusSucceeded StatusString = "succeeded"
)

type Status struct {
	StatusString StatusString
}

type Job struct {
	Id          string
	Name        string
	CurrentDate string
	Status      StatusString
	Repo        repo.Repo
}

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job, statusString StatusString) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
		status:     statusString,
	}
}

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
	status     StatusString
}

func (w Worker) start() {
	go func() {
		for {
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Start job
				job.Status = StatusRunning
				pkg.Info("worker%d: ID:%s started %s %s \n", w.id, job.Id, job.Status, job.Repo.Url)

				//Test Data Only
				Since := time.Date(2020, 12, 24, 11, 11, 53, 0, time.UTC)

				//Clone Git Repo.
				clone := repo.NewClone(job.Repo.Url, job.Repo.Branch, job.Repo.ConfigFile, Since)
				clone.ReadIntoMemory(w.id, job.Id)

				job.Status = StatusSucceeded
				pkg.Info("worker%d: ID:%s %s\n", w.id, job.Id, job.Status)
			case <-w.quitChan:
				//stop worker.
				pkg.Info("worker%d stopping\n", w.id)
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool, StatusQueued)
		worker.start()
	}
	go d.Dispatch()
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				// Create workers Job Queue
				workerJobQueue := <-d.workerPool
				pkg.Info("%s : ID:%s adding configuration %s\n", job.Status, job.Id, job.Name)
				//Dispatch a job to the Workers Job Queue
				workerJobQueue <- job
			}()
		}
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the delay.
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate delay is in range 1 to 10 seconds.
	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be between 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}

	// Set name and validate value.
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	// Create Job and push the work onto the jobQueue.
	job := Job{Name: name}
	jobQueue <- job

	// Render success.
	w.WriteHeader(http.StatusCreated)
}
