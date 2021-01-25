package viper

import "C"
import (
	"log"
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"github.com/spf13/viper"
	"time"
)

type AgcConfig struct {
	Name   string
	Url    string
	Branch string
	Config string
}

type Agc struct {
	//Apollo Guidance Configuration
	FileName string
	Directory string
}

func (a *Agc) Services() []AgcConfig {
	viper.SetConfigType("yaml")

	viper.SetConfigName(a.FileName)
	viper.AddConfigPath(a.Directory)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading the Apollo Guidance Configuration %s", err)
	}

	var services []AgcConfig
	err := viper.UnmarshalKey("repos", &services)
	if err != nil {
		panic("Unable to unmarshal hosts")
	}

	liftoff := pkg.FormatDate(time.Now())
	log.Printf("LIFT-OFF: %s\n", liftoff)

	return services
}

func SetAgc(fileName string, directory string) *Agc {

	return &Agc{
		FileName:      fileName,
		Directory:     directory,
	}
}