package types

import "time"

type Validator struct {
	OperatorAddress string          `json:"operator_address"`
	ConsensusPubkey ConsensusPubkey `json:"consensus_pubkey"`
	Description     Description     `json:"description"`
}

type ConsensusPubkey struct {
	Key string `json:"key"`
}

type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
}

type Participant struct {
	GitHub   string `json:"github"`
	PGP      string `json:"pgp"`
	Verified bool   `json:"verified"`
	Reason   string `json:"reason"`
	Addr     string `json:"addr"`
	ValAddr  string `json:"val_addr"`
	Tasks    Tasks  `json:"tasks"`
	Badge    Badge  `json:"badge"`
}

type Tasks struct {
	Task1 bool `json:"task1"`
	Task2 bool `json:"task2"`
	Task3 bool `json:"task3"`
}

type Badge struct {
	Silver int `json:"sliver"`
	Bronze int `json:"bronze"`
}

type GithubInfo struct {
	Username string `json:"username"`
}

type ServicesSummary struct {
	GithubInfo GithubInfo `json:"github"`
}

type KeybaseUser struct {
	ServicesSummary ServicesSummary `json:"services_summary"`
}
type KeybaseResponse struct {
	List []KeybaseUser `json:"list"`
}

type GithubResponse struct {
	CreatedTime time.Time `json:"created_at"`
}
