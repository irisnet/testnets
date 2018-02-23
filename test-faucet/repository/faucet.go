package repository

import "time"

type Faucet struct {
	Id      uint `gorm:"primary_key"`
	Address string
	Amount  int
	Coin    string
	Time    time.Time
}

func (faucet *Faucet) BeforeCreate() error {
	now := time.Now()
	faucet.Time = now
	return nil
}

func (faucet *Faucet) Create() error {
	return DB.Create(faucet).Error
}

func FindFaucet(address string, coin string) ([]Faucet, error) {
	var faucets []Faucet
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	err := DB.Where("address = ? and coin = ? and time >=? and time<?", address, coin, tm1, tm2).Find(&faucets).Error
	return faucets, err
}
