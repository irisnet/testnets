package main

import (
	"encoding/json"
	"io/ioutil"

	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/irisnet/irishub-sdk-go/types/store"

	"github.com/irisnet/testnets/bifrost/phase-2/count/tasks"
	counttypes "github.com/irisnet/testnets/bifrost/phase-2/count/types"
)

var (
	nodeURI  = "http://34.80.22.255:26657"
	grpcAddr = "34.80.22.255:9090"
	chainID  = "bifrost-2"
)

func main() {
	client := initClient()

	participants := getParticipants()

	tasks.CountTasks(client, participants)

	for _, participant := range participants {
		if participant.Tasks.Task1 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.Task2 {
			participant.Badge.Bronze += 1
		}
		if participant.Tasks.Task3 {
			participant.Badge.Silver += 1
		}
	}

	Silver := 0
	Bronze := 0

	for _, participant := range participants {
		Silver += participant.Badge.Silver
		Bronze += participant.Badge.Bronze
	}

	println("\nTotal Participants: ", len(participants))
	println("\nTotal Silver: ", Silver)
	println("\nTotal Bronze: ", Bronze)
	println()

	bz, _ := json.MarshalIndent(participants, "", "    ")
	if err := ioutil.WriteFile("result.json", bz, 0777); err != nil {
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
