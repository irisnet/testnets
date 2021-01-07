package main

import (
	"encoding/json"
	"io/ioutil"

	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/irisnet/irishub-sdk-go/types/store"

	"github.com/irisnet/testnets/bifrost/phase-1/count/tasks"
	counttypes "github.com/irisnet/testnets/bifrost/phase-1/count/types"
)

var (
	nodeURI  = "http://34.80.22.255:26657"
	grpcAddr = "34.80.22.255:9090"
	chainID  = "bifrost-1"
)

func main() {
	client := initClient()

	participants := getParticipants()

	tasks.CountServiceTasks(client, participants)
	tasks.CountRecordTasks(client, participants)
	tasks.CountNFTTasks(client, participants)
	tasks.CountRandomTasks(client, participants)

	for _, participant := range participants {
		if participant.Tasks.ServiceTasks.Task1 {
			participant.Badge.Silver += 1
		}
		if participant.Tasks.ServiceTasks.Task2 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.ServiceTasks.Task3 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.ServiceTasks.Task4 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.ServiceTasks.Task5 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.ServiceTasks.Task6 {
			participant.Badge.Gold += 1
		}
		if participant.Tasks.ServiceTasks.Task7 {
			participant.Badge.Silver += 1
		}
		if participant.Tasks.RecordTasks.Task1 {
			participant.Badge.Silver += 1
		}
		if participant.Tasks.NFTTasks.Task1 {
			participant.Badge.Silver += 1
		}
		if participant.Tasks.NFTTasks.Task2 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.NFTTasks.Task3 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.NFTTasks.Task4 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.NFTTasks.Task5 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.RandomTasks.Task1 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.RandomTasks.Task2 {
			participant.Badge.Bronze += 1
		}
	}

	Gold := 0
	Silver := 0
	Bronze := 0

	for _, participant := range participants {
		Gold += participant.Badge.Gold
		Silver += participant.Badge.Silver
		Bronze += participant.Badge.Bronze
	}

	println("Total: ", len(participants))
	println("Gold: ", Gold)
	println("Silver: ", Silver)
	println("Bronze: ", Bronze)

	bz, _ := json.MarshalIndent(participants, "", "    ")
	if err := ioutil.WriteFile("result.json", bz, 0666); err != nil {
		panic(err)
	}
}

func initClient() sdk.IRISHUBClient {
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
	}

	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	return sdk.NewIRISHUBClient(cfg)
}

func getParticipants() []*counttypes.Participant {
	bytes, err := ioutil.ReadFile("participants.json")
	if err != nil {
		panic(err)
	}

	var participants []*counttypes.Participant
	if err := json.Unmarshal(bytes, &participants); err != nil {
		panic(err)
	}

	return participants
}
