package tasks

import (
	"fmt"
	"strings"

	sdk "github.com/irisnet/irishub-sdk-go"
	types "github.com/irisnet/irishub-sdk-go/types"

	biftypes "github.com/irisnet/testnets/bifrost/phase-1/count/types"
)

func CountNFTTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	fmt.Println("count NFT task 1 ...")
	CountNFTTask1(client, participants)
	fmt.Println("count NFT task 2 ...")
	CountNFTTask2(client, participants)
	fmt.Println("count NFT task 3 ...")
	CountNFTTask3(client, participants)
	fmt.Println("count NFT task 4 and 5 ...")
	CountNFTTask4And5(client, participants)
}

func CountNFTTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("issue_denom")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			participant.Tasks.NFTTasks.Task1 = true
			bytes := []byte(participant.Tasks.Tasks)
			bytes[8] = '1'
			participant.Tasks.Tasks = string(bytes)
		}
	}
}

func CountNFTTask2(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("mint_nft")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 1 {
			participant.Tasks.NFTTasks.Task2 = true
			bytes := []byte(participant.Tasks.Tasks)
			bytes[9] = '1'
			participant.Tasks.Tasks = string(bytes)
		}
	}
}

func CountNFTTask3(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("edit_nft")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		for _, tx := range txs.Txs {
			uri, err := tx.Result.Events.GetValue("edit_nft", "token-uri")
			if err != nil {
				panic(err)
			}
			if strings.HasPrefix(uri, "bifrost-") {
				participant.Tasks.NFTTasks.Task3 = true
				bytes := []byte(participant.Tasks.Tasks)
				bytes[10] = '1'
				participant.Tasks.Tasks = string(bytes)
			}
		}
	}
}

func CountNFTTask4And5(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("transfer_nft")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			participant.Tasks.NFTTasks.Task4 = true
			bytes := []byte(participant.Tasks.Tasks)
			bytes[11] = '1'
			participant.Tasks.Tasks = string(bytes)
		}

		for _, tx := range txs.Txs {
			tokenID, err := tx.Result.Events.GetValue("transfer_nft", "token-id")
			if err != nil {
				panic(err)
			}
			builder := types.NewEventQueryBuilder().AddCondition(
				types.NewCond("message", "action").EQ(types.EventValue("burn_nft")),
			).AddCondition(
				types.NewCond("burn_nft", "token-id").EQ(types.EventValue(tokenID)),
			)

			txs, err := client.QueryTxs(builder, 1, 100)
			if err != nil {
				panic(err)
			}

			for _, tx := range txs.Txs {
				respTokenID, err := tx.Result.Events.GetValue("burn_nft", "token-id")
				if err != nil {
					panic(err)
				}
				if tokenID == respTokenID {
					participant.Tasks.NFTTasks.Task5 = true
					bytes := []byte(participant.Tasks.Tasks)
					bytes[12] = '1'
					participant.Tasks.Tasks = string(bytes)
				}
			}
		}
	}
}
