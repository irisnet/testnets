package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/irisnet/testnets/test-faucet/config"
)

func Addr(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"addr": config.Config.Iris})
}
