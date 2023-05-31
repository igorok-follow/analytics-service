package main

import (
	"flag"
	"github.com/igorok-follow/analytics-service/app"
	"github.com/igorok-follow/analytics-service/tools/location_generator"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/igorok-follow/analytics-service/config"
)

const defaultConfigPath = "./config/"

const defaultConfigName = "public_config.yaml"

var configName string

var configPath string

func init() {
	flag.StringVar(&configName, "config-name", defaultConfigName, "config name")
	flag.StringVar(&configPath, "config-path", defaultConfigPath, "config path")
}

func main() {
	flag.Parse()

	config, err := getConfig(configName)
	if err != nil {
		log.Printf("package main: config error \n%v", err)
	}

	location_generator.GenerateLocations("http://localhost:65001", config.Server.Name)

	app.Run(config)
}

func getConfig(name string) (*config.Config, error) {
	configPath = configPath + name

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config *config.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
