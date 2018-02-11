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
	resp, err := http.Get(server + "/query/nonce/" + iris)
	if err != nil {
		panic(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	print(string(body))
	var nonce Nonce
	json.Unmarshal(body, &nonce)
	nonce.Data += 1
	defer resp.Body.Close()

	//build send
	si := new(types.SendInput)
	si.Amount = types.Coins{types.Coin{Denom: denom, Amount: 10}}
	si.From = &types.Actor{ChainID: "", App: "sigs", Address: iris}
	si.To = &types.Actor{ChainID: "", App: "sigs", Address: addr}
	si.Fees = &types.Coin{Denom: denom, Amount: 1}
	si.Sequence = nonce.Data
	siStr, _ := json.Marshal(si)
	req, err := http.NewRequest("POST", server+"/build/send", bytes.NewBuffer([]byte(siStr)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)

	//sign tx
	requestSign := new(types.RequestSign)
	requestSign.Name = config.Config.Name
	requestSign.Password = config.Config.Password
	json.Unmarshal(body, &requestSign.Tx)
	rsStr, _ := json.Marshal(requestSign)
	req, err = http.NewRequest("POST", server+"/sign", bytes.NewBuffer([]byte(rsStr)))
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)

	//send tx
	req, err = http.NewRequest("POST", server+"/tx", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	if err == nil {
		faucet := &repository.Faucet{
			Address: addr,
			Amount:  amount,
		}
		err = faucet.Create()
	}
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
