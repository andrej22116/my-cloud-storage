package controllers

import (
	"database/sql"

	"./database"
)

// Глобальная переменная ( в рамках пакета )подключения к БД
var gDatabase *sql.DB

// Открывает соединение с БД PG по заданным параметрам
func OpenDatabaseConnection(host, port, dbname, user, password string) error {
	// Делаем конфиг
	var config database.Config
	// Заполняем содержимым
	config.ConnectionString = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	// Открываем соединение
	db, err := database.NewDatabase(&config)
	gDatabase = db

	return err
}
