package database

import (
	"github.com/MohammadMobasher/resturan-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySqlGormDB(config models.Configuration) *gorm.DB {
	dsn := config.MySqlUser + ":" + config.MySqlPassword + "@tcp(127.0.0.1:3306)/" + config.MySqlDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
