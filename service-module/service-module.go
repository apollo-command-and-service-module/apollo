package main

import (
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"github.com/apollo-command-and-service-module/apollo/pkg/job"
	cron "github.com/apollo-command-and-service-module/apollo/pkg/sync"
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, pkg.FormatDate(time.Now()))
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}


func main(){
	var (
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
		maxQueueSize = flag.Int("max_queue_size", 10, "The size of job queue")
		jobFrequency = flag.String("job_frequency", "@every 0h1m0s", "The start job frequency")
		port         = flag.String("port", "8080", "The server port")
	)
	flag.Parse()

	// Create the job queue.
	jobQueue := make(chan job.Job, *maxQueueSize)

	// Start the worker dispatcher.
	dispatcher := job.NewDispatcher(jobQueue, *maxWorkers)
	dispatcher.Run()

	// Start the Cron schedule
	fmt.Println(" Apollo, Houston. You're good at 1 minute.")
	cron.StartScheduler(jobFrequency, jobQueue)

	// http.HandleFunc("/healthcheck", healthcheck .Index)
	// Start the HTTP handler.
	// http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
	//	job.RequestHandler(w, r, jobQueue)
	// })
	setupRoutes()
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

