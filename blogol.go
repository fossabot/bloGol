package main

import (
	"log"

	"github.com/bloGol/bloGol/api"
	_ "github.com/bloGol/bloGol/init"
	"github.com/bloGol/bloGol/internal/db"
)

func main() {
	defer db.DB.Close()
	log.Println("Starting...")
	api.Run()
}
