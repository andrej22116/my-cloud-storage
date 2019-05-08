package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	ConnectionString string
}

type UserArguments struct {
	Login    string
	Password string
}

type UserData struct {
	AccessToken string
	Nickname    string
	Status      int8
}

func NewDatabase(config *Config) (*sql.DB, error) {
	connection, err := sql.Open("postgres", config.ConnectionString)

	if err != nil {
		return nil, err
	}
	if err = connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}
