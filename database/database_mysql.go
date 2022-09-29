package database

import (
	"database/sql"
	"time"

	"github.com/MohammadMobasher/resturan-backend/models"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySqlDB(config models.Configuration) *sql.DB {

	db, err := sql.Open("mysql", config.MySqlUser+":"+config.MySqlPassword+"@/"+config.MySqlDatabase)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
