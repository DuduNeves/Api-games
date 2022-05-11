package main

import (
	"github.com/DuduNeves/dev-games.git/database"
	"github.com/DuduNeves/dev-games.git/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
