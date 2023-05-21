package database

import (
	"class09/config"
	"class09/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	config.Init()

	db, err := gorm.Open(mysql.Open(config.Cfg.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Todo{}, &model.Customer{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
