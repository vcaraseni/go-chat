package main

import (
	"github.com/username/go-chat/internal/config"
	"github.com/username/go-chat/internal/database"
	"github.com/username/go-chat/internal/server"
)

func main() {
	config.loadConfig()
	database.ConnectDatabase()
	database.Migrate()

	server.StartServer(config.AppConfig.Server.Port)
}
