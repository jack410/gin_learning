package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Dsn       string
	SecretKey string
}

var Cfg *Config

func Init() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.db_name")

	// "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	Cfg = &Config{
		Dsn:       dsn,
		SecretKey: viper.GetString("SecretKey"),
	}
}

