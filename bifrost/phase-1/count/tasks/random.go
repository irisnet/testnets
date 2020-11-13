package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"

	sdk "github.com/irisnet/irishub-sdk-go"
	types "github.com/irisnet/irishub-sdk-go/types"

	biftypes "github.com/irisnet/testnets/bifrost/phase-1/count/types"
)

func CountRandomTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	fmt.Println("count Random task 1 and 2 ...")
	CountRandomTask1And2(client, participants)
}

func CountRandomTask1And2(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("request_rand")),
		)

		txs, err := client.QueryTxs(builder, 1, 10000)
		if err != nil {
			panic(err)
		}

		for _, tx := range txs.Txs {
			msgs := tx.Tx.GetMsgs()
			for _, msg := range msgs {
				bz, err := json.Marshal(msg)
				if err != nil {
					panic(err)
				}
				oracle := gjson.GetBytes(bz, "oracle").Bool()
				if !oracle {
					requestID, err := tx.Result.Events.GetValue("request_random", "request_id")
					if err != nil {
						panic(err)
					}
					if _, err = client.Random.QueryRandom(requestID); err == nil {
						participant.Tasks.RandomTasks.Task1 = true
					}
				} else {
					requestID, err := tx.Result.Events.GetValue("request_random", "request_id")
					if err != nil {
						panic(err)
					}
					if _, err = client.Random.QueryRandom(requestID); err == nil {
						participant.Tasks.RandomTasks.Task2 = true
					}
				}
			}
		}
	}
}
