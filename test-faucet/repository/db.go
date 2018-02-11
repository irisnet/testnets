package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/irisnet/testnets/test-faucet/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	postgres := config.Config.Postgres
	db, err := gorm.Open("postgres", "host=" + postgres.Host + " user=" + postgres.User+
		" dbname="+ postgres.Dbname+ " sslmode="+ postgres.Sslmode+ " password="+ postgres.Password)
	if err != nil {
		print(err)
	}

	//use singular table by default，else table name will add 's'
	db.SingularTable(true)

	//连接池上限
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(70)
	db.LogMode(true)
	DB = db

	//defer db.Close()
}
