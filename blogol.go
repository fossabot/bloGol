package main

import (
	"log"

	"github.com/bloGol/bloGol/api"
	_ "github.com/bloGol/bloGol/init"
	"github.com/bloGol/bloGol/internal/config"
	"github.com/bloGol/bloGol/internal/database"
)

func main() {
	defer database.DB.Close()
	log.Println("Starting...")
	api.API.Run(config.Config)
}
