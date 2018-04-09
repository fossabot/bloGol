package init

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bloGol/bloGol/api"
	"github.com/bloGol/bloGol/internal/config"
	"github.com/bloGol/bloGol/internal/database"
)

var pathPrefix = filepath.Join(
	os.Getenv("GOPATH"), "src", "github.com", "bloGol", "bloGol",
)

func init() {
	var err error
	config.Config, err = config.Open("./configs")
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("url:", config.Config.GetString("url"))
	log.Println("server.host:", config.Config.GetString("server.host"))
	log.Println("server.port:", config.Config.GetInt("server.port"))

	database.DB, err = database.Open(filepath.Join(pathPrefix, "test", "development.db"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	api.API = api.New()
}
