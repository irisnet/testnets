package types

type Participant struct {
	GitHub string   `json:"github"`
	PGP    string   `json:"pgp"`
	Addr   string   `json:"addr"`
	Serial uint     `json:"serial"`
	Tasks  [15]bool `json:"tasks"`
}

type Feed struct {
	FeedValues []Value `json:"feed_values"`
}

type Value struct {
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
}
