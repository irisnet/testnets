package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/irisnet/testnets/bifrost/phase-2/count/types"
)

const (
	KeybaseNotExist   = "Keybase Not Exist"
	GitHubNotExist    = "GitHub Not Exist"
	GitHubNotEligible = "GitHub Not Eligible"
)

var t, _ = time.ParseInLocation("2006-01-02 15:04:05", "2020-02-01 00:00:00", time.Local)

func main() {
	bytes, err := ioutil.ReadFile("participants_raw.json")
	if err != nil {
		panic(err)
	}

	var participants []types.Participant
	if err := json.Unmarshal(bytes, &participants); err != nil {
		panic(err)
	}

	for idx, participant := range participants {
		keybaseUrl := fmt.Sprintf("https://keybase.io/_/api/1.0/user/user_search.json?q=%s", participant.PGP)
		respKeybase, err := http.Get(keybaseUrl)
		if err != nil {
			panic(err)
		}
		defer respKeybase.Body.Close()

		bodyKeybase, err := ioutil.ReadAll(respKeybase.Body)
		if err != nil {
			panic(err)
		}

		var keybaseResp types.KeybaseResponse
		if err := json.Unmarshal(bodyKeybase, &keybaseResp); err != nil {
			panic(err)
		}

		if len(keybaseResp.List) == 0 {
			participants[idx].Reason = KeybaseNotExist
			println("x", "[", participant.GitHub, "]", KeybaseNotExist, participant.PGP)
			continue
		}

		githubUrl := fmt.Sprintf("https://api.github.com/users/%s", participant.GitHub)
		req, _ := http.NewRequest("GET", githubUrl, nil)
		req.Header.Set("Authorization", "token "+"4d9bfcd2917fc123f999dbc3ba5a2f05705d1585")
		respGitHub, err := (&http.Client{}).Do(req)
		if err != nil {
			panic(err)
		}
		defer respGitHub.Body.Close()

		if respGitHub.StatusCode != 200 {
			participants[idx].Reason = GitHubNotExist
			println("x", "[", participant.GitHub, "]", respGitHub.StatusCode)
			continue
		}

		bodyGitHub, err := ioutil.ReadAll(respGitHub.Body)
		if err != nil {
			panic(err)
		}

		var githubResp types.GithubResponse
		if err := json.Unmarshal(bodyGitHub, &githubResp); err != nil {
			panic(err)
		}

		if !githubResp.CreatedTime.Before(t) {
			participants[idx].Reason = GitHubNotEligible
			println("x", "[", participant.GitHub, "]", GitHubNotEligible)
			continue
		}

		participants[idx].Verified = true
		println("√", participant.GitHub)
	}

	bz, _ := json.MarshalIndent(participants, "", "	")
	if err = ioutil.WriteFile("participants_verified.json", bz, 0777); err != nil {
		panic(err)
	}
}
