package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/testnets/test-faucet/types"
	"os"
	"io"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

var iris = "0D7ACAD5C3F3EE3DBFB972F52D652509437E0044"
var server = "http://116.62.62.39:8999"

type TokenApply struct {
	Addr string
}

type Nonce struct {
	Height uint32
	Data   uint32
}

func apply(c *gin.Context) {
	var tokenApply TokenApply
	err := c.ShouldBindJSON(&tokenApply)
	if err != nil {
		return
	}
	addr := tokenApply.Addr
	resp, err := http.Get(server + "/query/nonce/" + iris)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	print(string(body))
	var nonce Nonce
	json.Unmarshal(body, &nonce)
	nonce.Data = +1
	defer resp.Body.Close()

	//build send
	si := new(types.SendInput)
	si.Amount = types.Coins{types.Coin{Denom: "iris", Amount: 10}}
	si.From = &types.Actor{ChainID: "", App: "sigs", Address: iris}
	si.To = &types.Actor{ChainID: "", App: "sigs", Address: addr}
	si.Fees = &types.Coin{Denom: "fermion", Amount: 1}
	si.Sequence = nonce.Data
	siStr, _ := json.Marshal(si)
	println(string(siStr))
	req, err := http.NewRequest("POST", server+"/build/send", bytes.NewBuffer([]byte(siStr)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	println(string(body))

	//sign tx
	requestSign := new(types.RequestSign)
	requestSign.Name = "iris"
	requestSign.Password = "1234567890"
	json.Unmarshal(body, &requestSign.Tx)
	rsStr, _ := json.Marshal(requestSign)
	println(string(rsStr))
	req, err = http.NewRequest("POST", server+"/byteTx", bytes.NewBuffer([]byte(rsStr)))
	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ = ioutil.ReadAll(resp.Body)
	println(string(body))
	println(addr)

}

func main() {
	r := gin.New()
	//log
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	log.SetOutput(gin.DefaultWriter) // You may need this

	r.POST("/apply", apply)
	r.Run("0.0.0.0:3434")
}
