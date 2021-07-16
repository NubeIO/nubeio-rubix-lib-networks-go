package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"rubix-lib-rest-go/config"
	"rubix-lib-rest-go/model"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

// Setup opens a database and saves the reference to `Database` struct.
func Setup() {
	var db = DB
	commonConfig := config.CommonConfig()
	driver := commonConfig.Database.Driver
	logs := commonConfig.Database.Logging

	if driver == "sqlite" {
		if logs {
			db, err = gorm.Open(sqlite.Open("./ugin.db?_foreign_keys=on"), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
		} else {
			db, err = gorm.Open(sqlite.Open("./ugin.db?_foreign_keys=on"), &gorm.Config{
			})
		}
		if err != nil {
			DBErr = err
			fmt.Println("db err: ", err)
		}
	}
	// Auto migrate project models
	err = db.AutoMigrate(&model.Network{}, &model.Device{}, &model.Point{}, &model.PointStore{})
	if err != nil {
		return
	}
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
// GetDBError helps you to get a connection
func GetDBError() error {
	return DBErr
}
