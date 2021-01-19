package cache_test

import (
	"fmt"
	"testing"

	"github.com/apollo-command-and-service-module/apollo/pkg/cache"
	"github.com/apollo-command-and-service-module/apollo/pkg/logging"
)

func TestConnectionRedis(t *testing.T) {
	log := logging.NewConsole(true)
	if err := cache.Genesis(); err != nil {
		log.PrintErrorf("Couldn't connect to Redis: %s", err)
	}

}

//Use Examples to check printout in string
func ExampleGetCache() {
	log := logging.NewConsole(true)
	jobId := 0
	c, err := cache.GetCache(jobId)
	if err != nil {
		log.PrintErrorf("Could not read from redis: %s", err)
	}
	fmt.Println(c.Title)

	// Output: redis genesis
}
