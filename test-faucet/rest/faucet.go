package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/testnets/test-faucet/types"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/irisnet/testnets/test-faucet/repository"
	"github.com/irisnet/testnets/test-faucet/config"
	"log"
)

type TokenApply struct {
	Addr string `binding:"required"`
}

type Nonce struct {
	Height uint32
	Data   uint32
}

func Apply(c *gin.Context) {
	var tokenApply TokenApply
	iris := config.Config.Iris
	server := config.Config.Client
	amount := config.Config.Amount
	denom := config.Config.Denom
	err := c.ShouldBindJSON(&tokenApply)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addr := tokenApply.Addr
	if !check(addr) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "已达当日上限"})
		return
	}
	result := doGet(server + "/query/nonce/" + iris)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var nonce Nonce
	json.Unmarshal(result, &nonce)
	nonce.Data += 1

	//build send
	si := new(types.SendInput)
	si.Amount = types.Coins{types.Coin{Denom: denom, Amount: 10}}
	si.From = &types.Actor{ChainID: "", App: "sigs", Address: iris}
	si.To = &types.Actor{ChainID: "", App: "sigs", Address: addr}
	si.Fees = &types.Coin{Denom: denom, Amount: 1}
	si.Sequence = nonce.Data
	siStr, _ := json.Marshal(si)
	result = doPost(server+"/build/send", siStr)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	//sign tx
	requestSign := new(types.RequestSign)
	requestSign.Name = config.Config.Name
	requestSign.Password = config.Config.Password
	json.Unmarshal(result, &requestSign.Tx)
	rsStr, _ := json.Marshal(requestSign)
	result = doPost(server+"/sign", rsStr)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	//send tx
	result = doPost(server+"/tx", result)
	if result == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	println(string(result))
	if err == nil {
		faucet := &repository.Faucet{
			Address: addr,
			Amount:  amount,
		}
		err = faucet.Create()
	}
	c.JSON(http.StatusOK, gin.H{"msg": "申请成功"})
}

func check(address string) bool {
	faucets, _ := repository.FindFaucetByAddress(address)
	var i int
	for _, faucet := range faucets {
		i = i + faucet.Amount
	}
	if i >= 100 {
		return false
	}
	return true
}

func doGet(url string) []byte {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func doPost(url string, data []byte) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
