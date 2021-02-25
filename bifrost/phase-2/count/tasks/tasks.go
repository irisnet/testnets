package tasks

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/irisnet/irishub-sdk-go"
	types "github.com/irisnet/irishub-sdk-go/types"

	biftypes "github.com/irisnet/testnets/bifrost/phase-2/count/types"
)

func CountTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	fmt.Println("\n\ncount task 1 ...")
	CountTask1(client, participants)
	fmt.Println("\n\ncount task 2 ...")
	CountTask2(client, participants)
	fmt.Println("\n\ncount task 3 ...")
	CountTask3(client, participants)
}

func CountTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		if !participant.Verified || participant.Addr == "" || participant.Tasks.Task1 {
			continue
		}

		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("proposal_vote", "proposal_id").EQ(types.EventValue(fmt.Sprintf("%d", 1))),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			println("√", participant.GitHub)
			participant.Tasks.Task1 = true
		}
	}
}

func CountTask2(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for _, participant := range participants {
		if !participant.Verified || participant.Addr == "" || participant.Tasks.Task2 {
			continue
		}

		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("proposal_vote", "proposal_id").EQ(types.EventValue(fmt.Sprintf("%d", 3))),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		if txs.Total > 0 {
			println("√", participant.GitHub)
			participant.Tasks.Task2 = true
		}
	}
}

func CountTask3(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	for i := int64(292989); i <= int64(293190); i += 1 {
		println("Height: ", i)
		resultBlock, _ := client.BaseClient.Block(context.Background(), &i)
		for _, signature := range resultBlock.Block.LastCommit.Signatures {
			for _, participant := range participants {
				if !participant.Verified || participant.Addr == "" || participant.Tasks.Task3 {
					continue
				}
				if participant.ValAddr == strings.ToUpper(hex.EncodeToString(signature.ValidatorAddress)) {
					println("√", participant.GitHub)
					participant.Tasks.Task3 = true
				}
			}
		}
	}
}
