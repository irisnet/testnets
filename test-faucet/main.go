package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"log"
	"github.com/irisnet/testnets/test-faucet/config"
	"github.com/irisnet/testnets/test-faucet/rest"
	"github.com/irisnet/testnets/test-faucet/repository"
	"github.com/gin-contrib/cors"
)

func main() {

	r := gin.New()
	//log
	gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter) // You may need this

	r.Use(cors.Default())
	r.POST("/apply", rest.Apply)
	r.GET("/addr", rest.Addr)

	repository.DB.Debug().AutoMigrate(&repository.Faucet{})

	r.Run(config.Config.Server)
}
