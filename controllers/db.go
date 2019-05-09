package controllers

import (
	"database/sql"

	"./database"
)

var gDatabase *sql.DB

func OpenDatabaseConnection() error {
	var config database.Config
	config.ConnectionString = "host=localhost port=5432 user=postgres password=1 dbname=cloud_storage sslmode=disable"
	db, err := database.NewDatabase(&config)
	gDatabase = db

	return err
}
