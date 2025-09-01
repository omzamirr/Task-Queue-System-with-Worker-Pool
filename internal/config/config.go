package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DATABASE_PATH string
	SERVER_PORT   string
	WORKER_COUNT  int
	ENVIRONMENT   string
}

func LoadConfig() (*Config, error) {
	config := &Config{}
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		config.DATABASE_PATH = "./data/taskqueue.db"
	} else {
		config.DATABASE_PATH = dbPath
	}

	serverPath := os.Getenv("SERVER_PORT")
	if serverPath == "" {
		config.SERVER_PORT = "8080"
	} else {
		config.SERVER_PORT = serverPath
	}

	envPath := os.Getenv("ENVIRONMENT")
	if envPath == "" {
		config.ENVIRONMENT = "development"
	} else {
		config.ENVIRONMENT = envPath
	}

	workPath := os.Getenv("WORKER_COUNT")
	if workPath == "" {
		config.WORKER_COUNT = 3
	} else {
		res, err := strconv.Atoi(workPath)
		if err != nil {
			return nil, fmt.Errorf("invalid WORKER_COUNT: %s", workPath)
		} else {
			config.WORKER_COUNT = res
		}

	}

	return config, nil

}
