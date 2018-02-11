package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
)

type Configuration struct {
	Postgres   *Postgres `yaml:"postgres"`
	Server     string    `yaml:"server"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Post     int    `yaml:"post"`
	Dbname   string `yaml:"dbname"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

type Redis struct {
	Url         string `yaml:"host"`
	ActcTimeout int    `yaml:"actcTimeout"`
	RescTimeout int    `yaml:"rescTimeout"`
	VercTimeOut int    `yaml:"vercTimeOut"`
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
