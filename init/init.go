package init

import (
	"log"

	"github.com/bloGol/bloGol/internal/config"
)

func init() {
	if err := config.Open("./configs"); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("url:", config.Config.GetString("url"))
	log.Println("server.host:", config.Config.GetString("server.host"))
	log.Println("server.port:", config.Config.GetInt("server.port"))
}
