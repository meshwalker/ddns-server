package ddns_server

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	"fmt"
)



func NewDB(dbhost, dbport, dbname, dbuser, dbpassword string) (*sql.DB,error){
	var pgconf string = "host=" + dbhost+ " "
	pgconf = pgconf + "port=" + dbport + " "
	pgconf = pgconf + "user=" + dbuser+ " "
	pgconf = pgconf + "password=" + dbpassword + " "
	pgconf = pgconf + "dbname=" + dbname + " "
	pgconf = pgconf + "sslmode=disable"


	fmt.Println(pgconf)
	db, err := sql.Open("postgres", pgconf)
	if err != nil {
		log.Fatal("Connection to postgres database failed")
	}


	erro := db.Ping()
	if erro != nil {
		log.Fatal("Can't ping")
	}

	if true {
		CreateTables(db)
	}

	return db, nil
}