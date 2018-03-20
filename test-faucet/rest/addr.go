package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/irisnet/testnets/test-faucet/config"
	"encoding/json"
)

type Account struct {
	Addr string
	Height uint32
	Data   Data
}

type Data struct {
	Coins []Coin
}

type Coin struct {
	Denom  string
	Amount uint32
}

func Addr(c *gin.Context) {
	server := config.Config.Client
	result := DoGet(server + "/query/account/" + config.Config.Iris)
	account := Account{Addr:config.Config.Iris}
	if result != nil {
		json.Unmarshal(result, &account)
	}

	c.JSON(http.StatusOK, account)
}
