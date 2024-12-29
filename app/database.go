package app

import (
	"book-rent-api/helper"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	configs := helper.GetConfigs()

	dsn := configs.Username+":@tcp("+configs.Host+":"+strconv.Itoa(configs.Port)+")/book_rent_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	
	if err != nil {
		panic(err)
	}
		
	sqlDB, err := db.DB()
		
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(5)

	return db
	
}