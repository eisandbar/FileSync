package utils

import (
	"os"
)

const (
	Server_HOST = "host"
	Server_PORT = "8000"
)

var Server_URI = "http://" + Server_HOST + ":" + Server_PORT

var (
	db_host = os.Getenv("DB_HOST")
	db_port = os.Getenv("DB_PORT")
)

var DB_URI = "http://" + db_host + ":" + db_port