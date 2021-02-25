package main

import (
	"encoding/json"
	counttypes "github.com/irisnet/testnets/bifrost/phase-2/count/types"
	"io/ioutil"
	"strconv"
)

func main() {
	participants := getResults()
	result := "| Github | PGP | Verified | Ineligible Description | Addr | Task1 | Task2 | Task3 | Sliver Badge | Bronze Badge |\n|:-------|:----|:---------|:-----------------------|:-----|:------|:------|:------|:-------------|:-------------|\n"
	var results1, results2, results3, results4 []*counttypes.Participant
	for _, participant := range participants{
		if participant.Verified && len(participant.Addr)>0{
			results1 = append(results1, participant)
		}else if participant.Verified {
			results2 = append(results2, participant)
		}else if len(participant.Addr)>0{
			results3 = append(results3, participant)
		}else {
			results4 = append(results4, participant)
		}
	}
	results1 = append(append(append(results1, results2...), results3...),results4...)
	for _, participant := range results1{
		github := participant.GitHub
		pgp := participant.PGP
		verified := strconv.FormatBool(participant.Verified)
		reason := participant.Reason
		addr := participant.Addr
		var task1, task2, task3, sliver, bronze string
		if participant.Verified {
			task1 = strconv.FormatBool(participant.Tasks.Task1)
			task2 = strconv.FormatBool(participant.Tasks.Task2)
			task3 = strconv.FormatBool(participant.Tasks.Task3)
			sliver = strconv.Itoa(participant.Badge.Silver)
			bronze = strconv.Itoa(participant.Badge.Bronze)
		}

		result += "| "+github+" | "+pgp+" | "+verified+" | "+reason+" | "+addr+" | "+task1+" | "+task2+" | "+task3+" | "+sliver+" | "+bronze+" |\n"
	}
	if err := ioutil.WriteFile("result.md", []byte(result), 0777); err != nil {
		panic(err)
	}
}

func getResults() []*counttypes.Participant {
	bytes, err := ioutil.ReadFile("result.json")
	if err != nil {
		panic(err)
	}

	var participants []*counttypes.Participant
	if err := json.Unmarshal(bytes, &participants); err != nil {
		panic(err)
	}

	return participants
}
