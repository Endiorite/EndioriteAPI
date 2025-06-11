package main

import (
	"EndioriteAPI/config"
	"EndioriteAPI/database"
	"os"
)

func main() {
	os.Setenv("TZ", "Europe/Paris")

	config.LoadEnv()
	database.ConnectMySQL()
	StartServer()
}
