package main

import (
	"EndioriteAPI/config"
	"EndioriteAPI/database"
)

func main() {
	config.LoadEnv()
	database.ConnectMySQL()
	StartServer()
}
