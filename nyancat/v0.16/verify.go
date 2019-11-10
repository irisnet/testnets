package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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
	moniker   string
	pgpID 	  string
	hash      string
	sender    string
	height    string
	timestamp string
}

type taskDetails []taskDetail						//实现sort接口，利用原生golang sort库
func (td taskDetails) Len() int {return len(td)}
func (td taskDetails) Swap(i,j int){td[i],td[j] = td[j],td[i]}
func (td taskDetails) Less(i,j int)bool{
	return td[i].timestamp < td[j].timestamp
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
func verifyTask12(url string) taskDetails {
	taskDetails := taskDetails{}
	data, err := fetch(url)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	txs := []interface{}{}
	totalPages,_ := strconv.Atoi(m["page_total"].(string))
	size,_ := strconv.Atoi(m["size"].(string))
	for i := 1; i <= totalPages;i++{
		url := fmt.Sprintf("%s&search_request_page=%d&search_request_size=%d",url,i,size)
		data, err := fetch(url)
		if err != nil {
			panic(err)
		}
		m := make(map[string]interface{})
		err = json.Unmarshal(data, &m)
		if err != nil {
			panic(err)
		}
		tx := m["txs"].([]interface{})
		txs = append(txs, tx...)
	}
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
			if validator,ok := validatorSet[sender];ok{
				detail.moniker = validator.moniker
				detail.pgpID = validator.pgp
			}
			//fmt.Println(detail)
			taskDetails = append(taskDetails, detail)
		}
	}
	return taskDetails
}


//task3 task4 任务检查，输入url以及token-pairs名，由于可能存在btc-eth,或者eth-btc均可，返回完成任务交易的相关情况
func verifyTask34(url, tokenPairs1, tokenPairs2 string) taskDetails {
	taskDetails := taskDetails{}
	data, err := fetch(url)
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	size,_  := strconv.Atoi(m["size"].(string))
	totalPages,_ := strconv.Atoi(m["page_total"].(string))
	txs := []interface{}{}
	for i := 1; i <= totalPages;i++{					//分页load，将所有信息拉取下来
		url := fmt.Sprintf("%s&search_request_page=%d&search_request_size=%d",url,i,size)
		data, err := fetch(url)
		if err != nil {
			panic(err)
		}
		m := make(map[string]interface{})
		err = json.Unmarshal(data, &m)
		if err != nil {
			panic(err)
		}
		tx := m["txs"].([]interface{})
		txs = append(txs, tx...)
	}
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
				if validator,ok := validatorSet[sender];ok{
					detail.moniker = validator.moniker
					detail.pgpID = validator.pgp
				}
				taskDetails = append(taskDetails, detail)
			}
		}
	}
	return taskDetails
}

func main() {
	task1Details := verifyTask12(task1URL)
	task2Details := verifyTask12(task2URL)
	task3Details := verifyTask34(task3URL, task3TP1, task3TP2)
	task4Details := verifyTask34(task4URL, task4TP1, task4TP2)
	sort.Sort(task1Details)
	removeRepeat(task1Details)
	uniSortTask1Details := removeRepeat(task1Details)
	fmt.Printf("task1完成人数统计：%d \n", len(uniSortTask1Details))
	for i,v := range uniSortTask1Details{
		fmt.Printf("【task1】%d: %+v\n",i,v)
	}

	sort.Sort(task2Details)
	removeRepeat(task2Details)
	uniSortTask2Details := removeRepeat(task2Details)
	fmt.Println()
	fmt.Printf("task2完成人数统计：%d \n", len(uniSortTask2Details))
	for i,v := range uniSortTask2Details{
		fmt.Printf("【task2】%d: %+v\n",i,v)
	}

	sort.Sort(task3Details)
	removeRepeat(task3Details)
	uniSortTask3Details := removeRepeat(task3Details)
	fmt.Println()
	fmt.Printf("task3完成人数统计：%d \n", len(uniSortTask3Details))
	for i,v := range uniSortTask3Details{
		fmt.Printf("【task3】%d: %+v\n",i,v)
	}

	sort.Sort(task4Details)
	removeRepeat(task4Details)
	uniSortTask4Details := removeRepeat(task4Details)
	fmt.Println()
	fmt.Printf("task4完成人数统计：%d \n", len(uniSortTask4Details))
	for i,v := range uniSortTask4Details{
		fmt.Printf("【task4】%d: %+v\n",i,v)
	}
}

func removeRepeat(details taskDetails)taskDetails{
	m := make(map[string]string)
	uniTaskDetails := taskDetails{}
	for _,detail := range details {
		if _,ok := m[detail.sender]; !ok{
			uniTaskDetails = append(uniTaskDetails, detail)
			m[detail.sender]="exit"
		}
	}
	return uniTaskDetails
}


//拉去网络资源
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
