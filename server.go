package ddns_server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/meshwalker/ddns-server/master/database"
)

func main() {
	cfg, err := GetConfig()
	if err != nil {
		log.Fatal("Can't read config values")
	}

	db, err := database.New(cfg.DbHost, cfg.Port, cfg.DbName, cfg.DbUser, cfg.DbPassword)
	if err != nil {
		log.Fatal("Can't connect to database!")
	}

	db.Ping()

	//_, err := db.Db.Query("SELECT * FROM domains")
}
