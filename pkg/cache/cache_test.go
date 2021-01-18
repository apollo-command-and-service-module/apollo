package cache_test

import (
	"fmt"
	"testing"

	"github.com/apollo-command-and-service-module/apollo/pkg/logging"

	"github.com/apollo-command-and-service-module/apollo/pkg/cache"
)

func TestConnectionRedis(t *testing.T) {

	log := logging.NewConsole(true)
	jobId := 0
	c, err := cache.GetCache(jobId)
	if err != nil {
		log.PrintError(err)
	}
	fmt.Println(c.Time)
}
