package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"

	sdk "github.com/irisnet/irishub-sdk-go"
	types "github.com/irisnet/irishub-sdk-go/types"

	biftypes "github.com/irisnet/testnets/bifrost/phase-1/count/types"
)

func CountServiceTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	fmt.Println("count Service task 1 ...")
	CountServiceTask1(client, participants)
	fmt.Println("count Service task 2 and 3 ...")
	CountServiceTask2And3(client, participants)
	fmt.Println("count Service task 4 and 5 ...")
	CountServiceTask4And5(client, participants)
	fmt.Println("count Service task 6 ...")
	CountServiceTask6(client, participants)
	fmt.Println("count Service task 7 ...")
	CountServiceTask7(client, participants)
}

func CountServiceTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("define_service")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			participant.Tasks.ServiceTasks.Task1 = true
		}
	}
}

func CountServiceTask2And3(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("bind_service")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
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
				serviceName := gjson.GetBytes(bz, "service_name").String()
				resp, err := client.Service.QueryServiceDefinition(serviceName)
				if err != nil {
					panic(err)
				}

				if resp.Author.String() == participant.Addr {
					participant.Tasks.ServiceTasks.Task2 = true
				} else {
					participant.Tasks.ServiceTasks.Task3 = true
				}
			}
		}
	}
}

func CountServiceTask4And5(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("call_service")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
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
				repeated := gjson.GetBytes(bz, "repeated").String()
				if repeated == "true" {
					participant.Tasks.ServiceTasks.Task5 = true
				} else {
					requestContextID, err := tx.Result.Events.GetValue("message", "request_context_id")
					if err != nil {
						panic(err)
					}

					block, err := client.QueryBlock(tx.Height)
					if err != nil {
						panic(err)
					}

					enblockRequestContextID, err := block.BlockResult.Results.EndBlock.Events.GetValue("new_batch_request", "request_context_id")
					if err == nil && requestContextID == enblockRequestContextID {
						participant.Tasks.ServiceTasks.Task4 = true
					}
				}
			}
		}
	}
}

func CountServiceTask6(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("respond_service")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			participant.Tasks.ServiceTasks.Task6 = true
		}
	}
}

func CountServiceTask7(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("withdraw_earned_fees")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			participant.Tasks.ServiceTasks.Task7 = true
		}
	}
}
