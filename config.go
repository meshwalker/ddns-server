package ddns_server

import (
	"os"
	"net/url"
	"strconv"
)

type Config struct {
	Url		*url.URL
	Port		string
	DbHost		string
	DbPort		string
	DbUser		string
	DbPassword	string
	DbName		string
	ClearDb		bool
}


func GetConfig() (*Config, error) {
	url, err := url.Parse(os.Getenv("URL"))
	if( err != nil ) {
		return nil, err
	}

	freshDb, err := strconv.ParseBool(os.Getenv("ClEAR_DB"))
	if( err != nil ) {
		return nil, err
	}

	cfg := Config{
		url,
		os.Getenv("PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		freshDb,
	}

	return &cfg, nil
}
