package main

import (
	"log"

	_ "github.com/bloGol/bloGol/init"
	"github.com/bloGol/bloGol/internal/db"
)

func main() {
	defer db.DB.Close()
	log.Println("Starting...")
}
