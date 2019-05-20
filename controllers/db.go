package controllers

import (
	"database/sql"

	"./database"
)

var gDatabase *sql.DB

func OpenDatabaseConnection(host, port, dbname, user, password string) error {
	var config database.Config
	config.ConnectionString = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := database.NewDatabase(&config)
	gDatabase = db

	return err
}
