package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var Config = readConfiguration()

type Repository struct {
	Url    string `yaml:"url"`
	Folder string `yaml:"folder"`
}

type Script struct {
	Location    string `yaml:"location"`
	Interpreter string `yaml:"interpreter"`
}

type Configuration struct {
	Repository  Repository `yaml:"repository"`
	Port        string     `yaml:"port"`
	Environment string     `yaml:"environment"`
	Script      Script     `yaml:"script"`
}

func readConfiguration() Configuration {
	file, err := ioutil.ReadFile("configuration.yml")

	if err != nil {
		log.Fatal(err)
	}

	configuration := Configuration{}

	err = yaml.Unmarshal(file, &configuration)

	if err != nil {
		log.Fatal(err)
	}
	return configuration
}
