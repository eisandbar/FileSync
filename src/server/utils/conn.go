package utils

import (
	"os"
	"fmt"
)

const (
	Server_HOST = "host"
	Server_PORT = "8000"
)

var Server_URI = "http://" + Server_HOST + ":" + Server_PORT

var (
	db_host = os.Getenv("DB_HOST")
	db_port = os.Getenv("DB_PORT")
	USER = os.Getenv("POSTGRES_USER")
	DB_NAME = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

var DB_URI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
	db_host,
	db_port,
	USER,
	DB_NAME,
	DB_PASSWORD,
)