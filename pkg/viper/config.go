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

func Services() []Repos {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// Apollo setting configurations
	var yamlExample = []byte(`
repos:
- name: Command-Module
  url: https://github.com/apollo-command-and-service-module/orbit.git
  branch: master
  config: config.yaml
- name: Lunar-Module
  url: https://github.com/apollo-command-and-service-module/orbit.git
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
