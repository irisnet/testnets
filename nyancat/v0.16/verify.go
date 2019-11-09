package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	task1URL = "https://lcd.nyancat.irisnet.org/txs?action=add_liquidity"
	task2URL = "https://lcd.nyancat.irisnet.org/txs?action=remove_liquidity"
	task3URL = "https://lcd.nyancat.irisnet.org/txs?action=swap_order"
	task4URL = "https://lcd.nyancat.irisnet.org/txs?action=swap_order"
	task3TP1 = "btc-iris" //task3的token-pairs对
	task3TP2 = "iris-btc"
	task4TP1 = "eth-btc" //task4的token-pairs对
	task4TP2 = "btc-eth"
)

type validator struct {				//用来将validator的pgp和ownkey进行匹配
	moniker string
	ownKey  string
	pgp     string
}

type taskDetail struct {			//每次任务查询后返回的原始结果
	hash      string
	sender    string
	height    string
	timestamp string
}

type taskStat struct {			//某个验证人四个任务是否完成
	task1 bool
	task2 bool
	task3 bool
	task4 bool
}

var validatorSet = make(map[string]validator)			//验证人集合
var taskTotal = make(map[string]taskStat)				//总的任务汇总
var notValidatorSender = make(map[string]taskStat)		//记录一些交易的发送者在验证人列表中找不到的信息
var repeatTxMap = make(map[string][]taskDetail)			//记录一些相同任务重复做的交易

//根据验证人的operationAddress从浏览器上面查找相应的ownerKey数据
func getOwnerKey(fva string) string {
	data, err := fetch("https://nyancat.irisplorer.io/api/stake/validators/" + fva)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	ownerKey := m["owner_addr"].(string)
	return ownerKey
}

//初始化是将验证人信息保存到本地的validatorSet集合
func init() {
	data, err := fetch("https://lcd.nyancat.irisnet.org/stake/validators")
	if err != nil {
		panic(err)
	}
	m := []interface{}{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	for _, v := range m {
		vali := validator{}
		validatorAddress := fmt.Sprint(v.(map[string]interface{})["operator_address"])
		moniker := fmt.Sprint(v.(map[string]interface{})["description"].(map[string]interface{})["moniker"])
		pgp := fmt.Sprint(v.(map[string]interface{})["description"].(map[string]interface{})["identity"])
		ownerAddress := getOwnerKey(validatorAddress)
		if _, ok := validatorSet[ownerAddress]; ok {
			fmt.Printf("%s 对应多个validator，请管理员检查\n", ownerAddress)
		} else {
			vali.moniker = moniker
			vali.pgp = pgp
			vali.ownKey = ownerAddress
			validatorSet[ownerAddress] = vali
		}
	}
}

//task1 task2的任务检查，输入url，返回为各个交易所需要的一些信息。
func verifyTask12(url string) []taskDetail {
	taskDetails := []taskDetail{}
	data, err := fetch(url)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	txs := m["txs"].([]interface{})

	for _, tx := range txs {
		code := tx.(map[string]interface{})["result"].(map[string]interface{})["Code"].(float64)
		if code != 0 {
			continue				//只记录成功的交易
		}
		hash := tx.(map[string]interface{})["hash"].(string)
		height := tx.(map[string]interface{})["height"].(string)
		timestamp := tx.(map[string]interface{})["timestamp"].(string)

		tags := tx.(map[string]interface{})["result"].(map[string]interface{})["Tags"].([]interface{})
		tagsMap := make(map[string]string)
		for _, tag := range tags {				//将tags解析保存到缓存中，去除json文件格式中的属性名
			key := tag.(map[string]interface{})["key"].(string)
			value := tag.(map[string]interface{})["value"].(string)
			tagsMap[key] = value
		}
		if sender, ok := tagsMap["sender"]; ok {
			detail := taskDetail{
				hash:      hash,
				height:    height,
				sender:    sender,
				timestamp: timestamp,				//获取交易的sender
			}
			//fmt.Println(detail)
			taskDetails = append(taskDetails, detail)
		}
	}
	return taskDetails
}


//task3 task4 任务检查，输入url以及token-pairs名，由于可能存在btc-eth,或者eth-btc均可，返回完成任务交易的相关情况
func verifyTask34(url, tokenPairs1, tokenPairs2 string) []taskDetail {
	taskDetails := []taskDetail{}
	data, err := fetch(url)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	txs := m["txs"].([]interface{})
	for _, tx := range txs {
		code := tx.(map[string]interface{})["result"].(map[string]interface{})["Code"].(float64)
		if code != 0 {		//只记录成功的交易
			continue
		}
		hash := tx.(map[string]interface{})["hash"].(string)
		height := tx.(map[string]interface{})["height"].(string)
		timestamp := tx.(map[string]interface{})["timestamp"].(string)

		tags := tx.(map[string]interface{})["result"].(map[string]interface{})["Tags"].([]interface{})
		tagsMap := make(map[string]string)
		for _, tag := range tags {
			key := tag.(map[string]interface{})["key"].(string)
			value := tag.(map[string]interface{})["value"].(string)
			tagsMap[key] = value
		}

		if tp, ok := tagsMap["token-pair"]; ok && (tp == tokenPairs1 || tp == tokenPairs2) {     //需要核对token-pairs对是否正确
			if sender, ok := tagsMap["sender"]; ok {
				detail := taskDetail{
					hash:      hash,
					height:    height,
					sender:    sender,
					timestamp: timestamp,
				}
				taskDetails = append(taskDetails, detail)
			}
		}
	}
	return taskDetails
}



func main() {

	task1Details := verifyTask12(task1URL)
	fmt.Printf("task1完成人数统计：%d \n", len(task1Details))
	task2Details := verifyTask12(task2URL)
	fmt.Printf("task2完成人数统计：%d \n", len(task2Details))
	task3Details := verifyTask34(task3URL, task3TP1, task3TP2)
	fmt.Printf("task3完成人数统计：%d \n", len(task3Details))
	task4Details := verifyTask34(task4URL, task4TP1, task4TP2)
	fmt.Printf("task4完成人数统计：%d \n", len(task4Details))

	senderMatchPGP(1, task1Details)
	senderMatchPGP(2, task2Details)
	senderMatchPGP(3, task3Details)
	senderMatchPGP(4, task4Details)

	fmt.Printf("从验证人人数统计：%d 人完成任务情况：\n", len(taskTotal))
	for k, v := range taskTotal {
		fmt.Printf("%s %+v \n\n", k, v)
	}

	fmt.Printf("未从验证人中找到人数统计：%d 人完成任务\n", len(notValidatorSender))
	for k, v := range notValidatorSender {
		fmt.Printf("%s %+v \n\n", k, v)
	}

	for k,v := range repeatTxMap {
		fmt.Printf("sender:%s 重复交易次数：%d\n",k,len(v))
	}

}

//将交易sender和PGP进行匹配
//如果能匹配上，则pgp：{true,true,false,true} ,其中分别代表task1，2，3，4是否完成
//如果匹配不上，则sender：{true,true,false,true}
func senderMatchPGP(taskID int, taskdetails []taskDetail) {
	repeatsender := []string{}
	for _, taskdetail := range taskdetails {    		//遍历每次任务结果
		ts := taskStat{}
		if validator, ok := validatorSet[taskdetail.sender]; ok {     //能够在validator集合中找到sender
			pgp := validator.pgp
			if _, ok := taskTotal[pgp]; ok {             //task2，3，4统计的时候需要用到task1中已经构建好的map
				repeatsender = append(repeatsender, taskdetail.sender)
				ts = taskTotal[pgp]
			}
			switch taskID {
			case 1:
				if !ts.task1 {
					ts.task1 = true
				} else {
					key := "TASK" + strconv.Itoa(taskID) + ":" + taskdetail.sender				//说明此sender已经发送过交易，本次交易重复了
					repeatTxArr := repeatTxMap[key]
					repeatTxArr = append(repeatTxArr, taskdetail)
					repeatTxMap[key] = repeatTxArr
				}

			case 2:
				if !ts.task2 {
					ts.task2 = true
				} else {
					key := "TASK" + strconv.Itoa(taskID) + ":" + taskdetail.sender
					repeatTxArr := repeatTxMap[key]
					repeatTxArr = append(repeatTxArr, taskdetail)
					repeatTxMap[key] = repeatTxArr
				}
			case 3:
				if !ts.task3 {
					ts.task3 = true
				} else {
					key := "TASK" + strconv.Itoa(taskID) + ":" + taskdetail.sender
					repeatTxArr := repeatTxMap[key]
					repeatTxArr = append(repeatTxArr, taskdetail)
					repeatTxMap[key] = repeatTxArr
				}
			case 4:
				if !ts.task4 {
					ts.task4 = true
				} else {
					key := "TASK" + strconv.Itoa(taskID) + ":" + taskdetail.sender
					repeatTxArr := repeatTxMap[key]
					repeatTxArr = append(repeatTxArr, taskdetail)
					repeatTxMap[key] = repeatTxArr
				}
			}
			taskTotal[pgp] = ts
		} else {
			if _, ok := notValidatorSender[taskdetail.sender]; ok {
				ts = notValidatorSender[taskdetail.sender]
			}
			switch taskID {
			case 1:
				ts.task1 = true
			case 2:
				ts.task2 = true
			case 3:
				ts.task3 = true
			case 4:
				ts.task4 = true
			}
			notValidatorSender[taskdetail.sender] = ts			//不再验证人中的找不到pgp，就用sender作为key
		}
	}
}


//拉取网络资源
func fetch(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
