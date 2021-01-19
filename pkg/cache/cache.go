package cache

import (
	"errors"
	"strconv"
	"time"

	"github.com/apollo-command-and-service-module/apollo/pkg/logging"

	"github.com/gomodule/redigo/redis"
)

//TODO:Consider passing this as an env variable or command-line option
var redisHost = "localhost:6379"
var pool *redis.Pool
var log *logging.Logger
var ErrNoEntry = errors.New("no entry found or job deleted")

//Since struct defines the cache data structure
type Since struct {
	Job   int    `redis:"job"`
	Title string `redis:"title"`
	Time  string `redis:"time"`
}

func init() {

	log = logging.NewConsole(true)

	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisHost) //TODO: needs to be changed to a static-ip addres or service discovery
		},
	}

}

//TODO: we should consider whether this should be exported. For now leaving as exported, but perhaps
//needs to be moved to our API as a maintenace option.

//SeedCache is used for populating the Redis cache on first run
func SeedCache() error {
	var job int
	since := &Since{
		Job:   job,
		Title: "redis seed",
		Time:  time.Now().String(),
	}
	conn := pool.Get()
	defer conn.Close()

	jobNumber := strconv.Itoa(since.Job)
	_, err := conn.Do("HSET", "job:"+jobNumber, "title", since.Title, "time", since.Time)
	if err != nil {
		log.PrintErrorf("An error occurred: %s", err)
		return err
	}

	return nil
}

//ReadCache reads from the redis cache based on job number
func ReadCache(job int) (*Since, error) {
	jobNumber := strconv.Itoa(job)
	conn := pool.Get()
	defer conn.Close()

	values, err := redis.StringMap(conn.Do("HGETALL", "job:"+jobNumber))
	if err != nil {
		log.PrintErrorf("Could not read from redis cache: %s", err)
		return nil, err
	}

	return &Since{
		Job:   job,
		Title: values["title"],
		Time:  values["time"],
	}, nil
}

//WriteCache writes to cache by passing in the job number/id and any title/info that we want to add
func WriteCache(job int, title string) (*Since, error) {
	jobNumber := strconv.Itoa(job)
	conn := pool.Get()
	defer conn.Close()

	since := &Since{
		Job:   job,
		Title: title,
		Time:  time.Now().String(),
	}

	_, err := conn.Do("HSET", "job:"+jobNumber, "title", since.Title, "time", since.Time)
	if err != nil {
		log.PrintErrorf("Could not write to redis cache: %s", err)
		return nil, err
	}

	return since, nil
}

//DeleteJobInCache removes the job in the redis cache
func DeleteJobInCache(job int) error {
	jobNumber := strconv.Itoa(job)
	conn := pool.Get()
	defer conn.Close()

	//check if the given job number exists
	exists, err := redis.Int(conn.Do("DEL", "job:"+jobNumber))
	if err != nil {
		log.PrintErrorf("Could not read from redis cache: %s", err)
		return err
	} else if exists == 0 {
		log.PrintInfof("Job %s doesn't exist: %s", jobNumber, ErrNoEntry)
		return ErrNoEntry
	}

	return nil
}

//TODO: This should be moved to a maintenance route, or called only from a maintenance API

//ClearRedisCacheAll clears the entire cache.
func ClearRedisCacheAll() error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")
	if err != nil {
		log.PrintErrorf("Could not flush redis cache %s", err)
		return err
	}

	return nil
}

//TODO: set client name in REDIS
//TODO: consider separate clients wanting to write to redis, we can then change the database key (max 15 by default available in redis)
