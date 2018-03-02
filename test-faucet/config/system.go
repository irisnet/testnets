package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Configuration struct {
	Postgres Postgres `yaml:"postgres"`
	Server   string    `yaml:"server"`
	Total    int       `yaml:"total"`
	Amount   int       `yaml:"amount"`
	Coin    string    `yaml:"coin"`
	Iris     string    `yaml:"iris"`
	Name     string    `yaml:"name"`
	Password string    `yaml:"password"`
	Client   string    `yaml:"client"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Post     int    `yaml:"post"`
	Dbname   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

var Config *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	//var config Configuration
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	//Config = &config
	log.Printf("config load succeessfully:%v", Config)
	return nil
}
