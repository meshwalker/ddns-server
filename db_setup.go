package ddns_server

import (
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
)


func CreateTables(db *sql.DB) {
	var autoUpdateDateTrigger string =
		"CREATE OR REPLACE FUNCTION auto_update_timestamp()"+
		"RETURNS TRIGGER AS $$"+
		"BEGIN"+
		"NEW.modified = now();"+
		"RETURN NEW;"+
		"END;"+
		"$$ language 'plpgsql';"

	var subDomainsModifiedTrigger string = "CREATE TRIGGER update_subdomains_modtime BEFORE UPDATE ON subdomains FOR EACH ROW EXECUTE PROCEDURE  auto_update_timestamp();"
	var recordsModifiedTrigger string = "CREATE TRIGGER update_records_modtime BEFORE UPDATE ON records FOR EACH ROW EXECUTE PROCEDURE  auto_update_timestamp();"


	var domainsTable string = "CREATE TABLE domains (" +
		"id SERIAL PRIMARY KEY NOT NULL," +
		"name VARCHAR(100) NOT NULL );"

	var subDomainsTable string = "CREATE TABLE subdomains (" +
		"id BIGSERIAL PRIMARY KEY NOT NULL," +
		"name VARCHAR(256) NOT NULL,"+
		"user_id bigint NOT NULL,"+
		"modified TIMESTAMP );"

	var recordsTable string = "CREATE TABLE records (" +
		"id BIGSERIAL PRIMARY KEY NOT NULL," +
		"subdomains_id bigint NOT NULL,"+
		"name VARCHAR(256) NOT NULL, "+
		"inet VARCHAR(256) NOT NULL, "+
		"type CHAR(8) NOT NULL, "+
		"ttl INTEGER NOT NULL,"+
		"modified TIMESTAMP,"+
		"FOREIGN KEY (subdomains_id) REFERENCES subdomains (id) ));"

	var err error

	_ , err = db.Query(domainsTable)
	if err != nil {
		log.Error("Can't create table: \"domains\"")
	} else {
		log.Info("Successfully created table: \"domains\"")
	}


	_ , err = db.Query(subDomainsTable)
	if err != nil {
		log.Error("Can't create table: \"subdomains\"")
	} else {
		log.Info("Successfully created table: \"subdomains\"")
	}


	_ , err = db.Query(recordsTable)
	if err != nil {
		log.Error("Can't create table: \"records\"")
	} else {
		log.Info("Successfully created table: \"records\"")
	}


	//Auto update field "modified" when update gets called
	_ , err = db.Query(autoUpdateDateTrigger)
	if err != nil {
		log.Error("Can't create function")
	} else {
		log.Info("Successfully created function")
	}


	_ , err = db.Query(subDomainsModifiedTrigger)
	if err != nil {
		log.Error("Can't create subDomainsModifiedTrigger")
	} else {
		log.Info("Successfully created subDomainsModifiedTrigger")
	}

	_ , err = db.Query(recordsModifiedTrigger)
	if err != nil {
		log.Error("Can't create recordsModifiedTrigger")
	} else {
		log.Info("Successfully created recordsModifiedTrigger")
	}
}
