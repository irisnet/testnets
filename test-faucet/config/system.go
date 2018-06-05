package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Configuration struct {
	Postgres Postgres `yaml:"postgres"`
	Server   string   `yaml:"server"`
	Total    int      `yaml:"total"`
	Amount   int      `yaml:"amount"`
	Coin     string   `yaml:"coin"`
	Iris     string   `yaml:"iris"`
	Name     string   `yaml:"name"`
	Password string   `yaml:"password"`
	Client   string   `yaml:"client"`
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

func init() {
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Panic(err.Error())
	}
	//var config Configuration
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Panic(err.Error())
	}
	//Config = &config
	log.Printf("config load succeessfully:%v", Config)

	//modify config bu env
	postgresHost := os.Getenv("DB_HOST")
	if postgresHost != "" {
		Config.Postgres.Host = postgresHost
	}
	password := os.Getenv("PASSWORD")
	if password != "" {
		Config.Password = password
	}

	if Config.Password == "" {
		log.Panic("faucet key password is empty")
	}
}
