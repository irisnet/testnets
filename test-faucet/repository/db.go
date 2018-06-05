package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/irisnet/testnets/test-faucet/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"log"
)

var DB *gorm.DB

func init() {
	db, err := initDB()
	if err != nil {
		log.Println("failed to init db,retry after 30s")
		time.Sleep(30 * time.Second)
		db, err = initDB()
		if err != nil {
			log.Panic(err.Error())
		}
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

func initDB() (*gorm.DB, error) {
	postgres := config.Config.Postgres
	return gorm.Open("postgres", "host=" + postgres.Host + " user=" + postgres.User+
		" dbname="+ postgres.Dbname+ " sslmode="+ postgres.Sslmode+ " password="+ postgres.Password)
}

