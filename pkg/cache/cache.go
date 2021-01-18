package cache

import (
	"strconv"
	"time"

	"github.com/apollo-command-and-service-module/apollo/pkg/logging"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

var job int

type Since struct {
	Job   int    `redis:"job"`
	Title string `redis:"title"`
	Time  string `redis:"time"`
}

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379") //TODO: needs to be changed to a static-ip addres or service discovery
		},
	}

	since := &Since{
		Job:   job,
		Title: "redis genesis",
		Time:  time.Now().String(),
	}
	conn := pool.Get()
	defer conn.Close()
	jobNumber := strconv.Itoa(since.Job)
	_, err := conn.Do("HSET", "job:"+jobNumber, "title", since.Title, "time", since.Time)
	if err != nil {
		log := logging.NewConsole(true)
		log.PrintErrorf("An error occurred: %s", err)
	}
}

func GetCache(job int) (*Since, error) {
	jobNumber := strconv.Itoa(job)
	conn := pool.Get()
	defer conn.Close()
	log := logging.NewConsole(true)

	values, err := redis.StringMap(conn.Do("HGETALL", "job:"+jobNumber))
	if err != nil {
		log.PrintErrorf("Could not read from redis: %s", err)
		return nil, err
	}

	return &Since{
		Job:   job,
		Title: values["title"],
		Time:  values["time"],
	}, nil
}
