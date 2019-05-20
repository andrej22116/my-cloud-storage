package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

/*
Config - Конфигурационная структура - для будущих расширений
*/
type Config struct {
	ConnectionString string
}

/*
UserArguments - параметры пользователя для авторизации
*/
type UserArguments struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

/*
UserData - Структура с информацией о пользователе (Status - для будущих улучшений)
*/
type UserData struct {
	AccessToken string `json:"token"`
	Nickname    string `json:"nickname"`
	Status      int8   `json:"status"`
}

/*
Token - токен пользователя. Используется для десереализации
*/
type Token struct {
	Token string `json:"token"`
}

/*
NewDatabase - создаёт новое подключение к БД
*/
func NewDatabase(config *Config) (*sql.DB, error) {
	// Подключаемся
	connection, err := sql.Open("postgres", config.ConnectionString)

	// Если ошибки - кидаем их
	if err != nil {
		return nil, err
	}
	// Пингуем сервер
	if err = connection.Ping(); err != nil {
		return nil, err
	}

	// Возвращаем соединение
	return connection, nil
}
