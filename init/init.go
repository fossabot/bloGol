package init

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bloGol/bloGol/internal/config"
	"github.com/bloGol/bloGol/internal/db"
)

var pathPrefix = filepath.Join(
	os.Getenv("GOPATH"), "src", "github.com", "bloGol", "bloGol",
)

func init() {
	if err := config.Open("./configs"); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("url:", config.Config.GetString("url"))
	log.Println("server.host:", config.Config.GetString("server.host"))
	log.Println("server.port:", config.Config.GetInt("server.port"))
	if err := db.Open(filepath.Join(pathPrefix, "test", "development.db")); err != nil {
		log.Fatalln(err.Error())
	}
}
