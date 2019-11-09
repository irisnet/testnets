package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

type validator struct {
	moniker          string
	operationAddress string
	pgp              string
}

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

var validatorSet = make(map[string]validator)

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
			vali.operationAddress = ownerAddress
			validatorSet[ownerAddress] = vali
		}
	}
}

type taskDetail struct {
	sender    string
	hash      string
	height    string
	timestamp string
}

func verifyTask12(url string) []taskDetail {
	taskDetails := []taskDetail{}
	data, err := fetch("https://lcd.nyancat.irisnet.org/txs?action=add_liquidity")
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
			continue
		}

		hash := tx.(map[string]interface{})["hash"].(string)
		height := tx.(map[string]interface{})["height"].(string)
		timestamp := tx.(map[string]interface{})["timestamp"].(string)

		//tags := tx.(map[string]interface{})["result"].(map[string]interface{})["Tags"].([]interface{})
		tags := tx.(map[string]interface{})["result"].(map[string]interface{})["Tags"].([]interface{})
		tagsMap := make(map[string]string)
		for _, tag := range tags {
			key := tag.(map[string]interface{})["key"].(string)
			value := tag.(map[string]interface{})["value"].(string)
			tagsMap[key] = value
		}
		if sender, ok := tagsMap["sender"]; ok {
			detail := taskDetail{
				hash:      hash,
				height:    height,
				sender:    sender,
				timestamp: timestamp,
			}
			//fmt.Println(detail)
			taskDetails = append(taskDetails, detail)
		}
	}
	return taskDetails
}

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
		if code != 0 {
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

		if tp, ok := tagsMap["token-pair"]; ok && (tp == tokenPairs1 || tp == tokenPairs2) {
			if sender, ok := tagsMap["sender"]; ok {
				detail := taskDetail{
					hash:      hash,
					height:    height,
					sender:    sender,
					timestamp: timestamp,
				}
				//fmt.Println(detail)
				taskDetails = append(taskDetails, detail)
			}
		}
	}

	return taskDetails

}

type taskStat struct {
	task1 bool
	task2 bool
	task3 bool
	task4 bool
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
	taskTotal := make(map[string]taskStat)
	notValidatorSender := make(map[string]taskStat)

	senderMatchPGP(1, task1Details, taskTotal, notValidatorSender)
	senderMatchPGP(2, task2Details, taskTotal, notValidatorSender)
	senderMatchPGP(3, task3Details, taskTotal, notValidatorSender)
	senderMatchPGP(4, task4Details, taskTotal, notValidatorSender)

	fmt.Printf("验证人人数统计：%d 人完成任务\n", len(taskTotal))
	for k, v := range taskTotal {
		fmt.Printf("%s %+v \n\n", k, v)
	}

	fmt.Printf("为从验证人找到人数统计：%d 人完成任务\n", len(notValidatorSender))
	for k, v := range notValidatorSender {
		fmt.Printf("%s %+v \n\n", k, v)
	}

}

func senderMatchPGP(taskID int, taskdetails []taskDetail, taskTotal map[string]taskStat, notValidatorSender map[string]taskStat) {
	i := 1
	okCount := 1
	failOunt := 1
	for _, taskdetail := range taskdetails {
		fmt.Printf("task%d:%d times!\n", taskID, i)
		i++
		ts := taskStat{}
		if validator, ok := validatorSet[taskdetail.sender]; ok {
			fmt.Printf("task%d:%d OKtimes!\n", taskID, okCount)
			okCount++

			pgp := validator.pgp
			if _, ok := taskTotal[pgp]; ok {
				ts = taskTotal[pgp]
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
			taskTotal[pgp] = ts
		} else {
			fmt.Printf("task%d:%d FAILtimes!\n", taskID, failOunt)
			failOunt++
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
			notValidatorSender[taskdetail.sender] = ts
		}
	}
}
