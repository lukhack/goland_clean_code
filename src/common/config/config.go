package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type ConfigDB struct {
	Host     string
	UserName string
	Password string
	Database string
	Port     int
	ChartSet string
}

type ConfigApp struct {
	Port string
}

type Config struct {
	App *ConfigApp
	DB  *ConfigDB
}

const BAD_REQUES = "BAD_REQUEST: %v"

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed to load the .env file")
		panic(err)
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Config{
		&ConfigApp{Port: os.Getenv("APP_PORT")},
		&ConfigDB{
			Host:     os.Getenv("DB_HOST"),
			UserName: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			Port:     port,
			ChartSet: os.Getenv("DB_CHARSET")},
	}
}
