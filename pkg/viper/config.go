package viper

import "C"
import (
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"bytes"
	"github.com/spf13/viper"
	"time"
)

type Repos struct {
	Name string
	Url string
	Branch string
	Config string
}


//TODO: find a location for apollo configuration / setting
// Do we want to store the config in S3 or a github repos that we pass environment variables

func Services() []Repos {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// Apollo setting configurations
	var yamlExample = []byte(`
repos:
- name: Command-Module
  url: https://github.com/apollo-command-and-service-module/orbit.git
  token: auth to get into github
  branch: master
  config: config.yaml
- name: gocd
  url: https://github.com/gocd/gocd.git
  branch: master
  config: config.yaml
- name: Service-Module
  url: https://github.com/apollo-command-and-service-module/orbit.git
  branch: master
  config: config.yaml
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	var services []Repos
	err := viper.UnmarshalKey("repos", &services)
	if err != nil {
		panic("Unable to unmarshal hosts")
	}

	liftoff:= pkg.FormatDate(time.Now())
	pkg.Info("LIFT-OFF: %s\n", liftoff)

	return services
}
