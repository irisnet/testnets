package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/modules/record"
	types "github.com/irisnet/irishub-sdk-go/types"

	biftypes "github.com/irisnet/testnets/bifrost/phase-1/count/types"
)

func CountRecordTasks(client sdk.IRISHUBClient, participants []*biftypes.Participant) {
	fmt.Println("count Record task 1 ...")
	CountRecordTask1(client, participants)
}

func CountRecordTask1(client sdk.IRISHUBClient, participants []*biftypes.Participant) {

	feed := getFeedValues()

	for _, participant := range participants {
		builder := types.NewEventQueryBuilder().AddCondition(
			types.NewCond("message", "sender").EQ(types.EventValue(participant.Addr)),
		).AddCondition(
			types.NewCond("message", "action").EQ(types.EventValue("create_record")),
		)

		txs, err := client.QueryTxs(builder, 1, 100)
		if err != nil {
			panic(err)
		}

		for _, tx := range txs.Txs {
			recordID, err := tx.Result.Events.GetValue("create_record", "record_id")
			if err != nil {
				panic(err)
			}

			request := record.QueryRecordReq{
				RecordID: recordID,
			}

			resp, err := client.Record.QueryRecord(request)
			if err != nil {
				panic(err)
			}

			if resp.Record.Contents[0].Digest == feed.FeedValues[participant.Serial-1].Data {
				participant.Tasks.RecordTasks.Task1 = true
				bytes := []byte(participant.Tasks.Tasks)
				bytes[7] = '1'
				participant.Tasks.Tasks = string(bytes)
			}
		}
	}
}

func getFeedValues() biftypes.Feed {
	bytes, err := ioutil.ReadFile("feed_values.json")
	if err != nil {
		panic(err)
	}

	var feed biftypes.Feed
	if err := json.Unmarshal(bytes, &feed); err != nil {
		panic(err)
	}

	return feed
}
