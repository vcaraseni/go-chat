package main

import (
	"github.com/vcaraseni/go-chat/internal/config"
	"github.com/vcaraseni/go-chat/internal/database"
	"github.com/vcaraseni/go-chat/internal/server"
)

func main() {
	config.LoadConfig()
	database.ConnectDatabase()
	database.Migrate()

	server.StartServer(config.AppConfig.Server.Port)
}
