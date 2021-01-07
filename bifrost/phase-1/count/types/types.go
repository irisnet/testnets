package types

type Participant struct {
	GitHub string `json:"github"`
	PGP    string `json:"pgp"`
	Addr   string `json:"addr"`
	Serial uint   `json:"serial"`
	Tasks  Tasks  `json:"tasks"`
	Badge  Badge  `json:"badge"`
}

type Tasks struct {
	ServiceTasks ServiceTasks `json:"service_tasks"`
	RecordTasks  RecordTasks  `json:"record_tasks"`
	NFTTasks     NFTTasks     `json:"nft_tasks"`
	RandomTasks  RandomTasks  `json:"random_tasks"`
	Tasks        string       `json:"tasks"`
}

type ServiceTasks struct {
	Task1 bool `json:"task1"`
	Task2 bool `json:"task2"`
	Task3 bool `json:"task3"`
	Task4 bool `json:"task4"`
	Task5 bool `json:"task5"`
	Task6 bool `json:"task6"`
	Task7 bool `json:"task7"`
}

type RecordTasks struct {
	Task1 bool `json:"task1"`
}

type NFTTasks struct {
	Task1 bool `json:"task1"`
	Task2 bool `json:"task2"`
	Task3 bool `json:"task3"`
	Task4 bool `json:"task4"`
	Task5 bool `json:"task5"`
}

type RandomTasks struct {
	Task1 bool `json:"task1"`
	Task2 bool `json:"task2"`
}

type Feed struct {
	FeedValues []Value `json:"feed_values"`
}

type Value struct {
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
}

type Badge struct {
	Gold   int `json:"gold"`
	Silver int `json:"sliver"`
	Bronze int `json:"bronze"`
}
