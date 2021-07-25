package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
)

type DatabaseConfig struct {
	User           string `yaml:"User"`
	Password       string `yaml:"Password"`
	Host           string `yaml:"Host"`
	Port           string `yaml:"Port"`
	DBName         string `yaml:"DBName"`
	CollectionName string `yaml:"CollectionName"`
}

type ServerConfig struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:server`
}

func GetConfig() (Config, error) {

	filename, _ := filepath.Abs("./dev_config.yaml")

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return config, nil
}
