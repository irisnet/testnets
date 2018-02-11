package repository

import "time"

type Faucet struct {
	Id      uint `gorm:"primary_key"`
	Address string
	Amount  int
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
