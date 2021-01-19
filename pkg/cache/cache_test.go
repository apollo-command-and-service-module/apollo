package cache_test

import (
	"testing"

	"github.com/apollo-command-and-service-module/apollo/pkg/cache"
	"github.com/apollo-command-and-service-module/apollo/pkg/logging"
	"github.com/gomodule/redigo/redis"
)

var log *logging.Logger
var redisHost = "localhost:6379"

func init() {
	log = logging.NewConsole(true)
}

func createSeed() error {

	conn, err := redis.Dial("tcp", redisHost)
	//check if seed exists
	exists, err := redis.Int(conn.Do("EXISTS", "job:0"))
	if err != nil {
		return err
	} else if exists == 0 {
		log.PrintInfof("Testing: Redis not populated, populating")
		err := cache.SeedCache()
		if err != nil {
			return err
		}
	}

	return nil
}

func setup(job int, title string) error {

	_, err := cache.WriteCache(job, title)

	if err != nil {
		return err
	}
	return nil

}

func TestReadFromCache(t *testing.T) {
	jobId := 0
	if err := createSeed(); err != nil {
		t.Logf("Error running test: %s", err)
		t.Fail()
	}
	c, err := cache.ReadCache(jobId)
	if err != nil {
		log.PrintErrorf("Could not read from redis cache: %s", err)
	}
	got := c.Title
	want := "redis seed"

	if got != want {
		t.Fail()
		t.Logf("Got %s; wanted: %s", got, want)
	}
}

func TestWriteToCache(t *testing.T) {
	jobId := 1
	title := "hello world"
	_, err := cache.WriteCache(jobId, title)
	if err != nil {
		log.PrintErrorf("Could not write to redis cache: %s", err)
	}
}

func TestDeleteFromCache(t *testing.T) {
	jobId := 1
	want := ""
	err := cache.DeleteJobInCache(jobId)
	if err != nil {
		log.PrintErrorf("Error occurred: %s", err)
		t.Logf("Got: %s; want: %s", err, want)
	}

}
