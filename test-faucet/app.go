package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"log"
	"github.com/irisnet/testnets/test-faucet/config"
	"github.com/irisnet/testnets/test-faucet/rest"
	"github.com/irisnet/testnets/test-faucet/repository"
)

func main() {

	//init config
	if err := config.LoadConfiguration("./config.yml"); err != nil {
		log.Print("config error")
		return
	}

	//db
	repository.InitDB()

	r := gin.New()
	//log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter) // You may need this

	r.POST("/apply", rest.Apply)
	r.Run(config.Config.Server)
}
