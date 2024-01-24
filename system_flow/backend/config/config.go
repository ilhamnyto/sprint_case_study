package config 

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

const (
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
)

func LoadConfig(filename string) {
	godotenv.Load(filename)
}

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) int {
	str := os.Getenv(key)

	num, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return num
}