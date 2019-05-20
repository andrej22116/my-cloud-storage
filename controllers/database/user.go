package database

import (
	"database/sql"
)

/*
Константы с запросами
*/
const (
	registrationNewUser = "select from registration_new_user($1, $2);"
	authorizationUser   = "select * from authorization_user($1, $2);"
	testUserSessionKey  = "select * from test_user_session_key($1);"
	logoutUser          = "select from logout_user($1);"
)

/*
Registration - запрос на регистрацию пользователя
*/
func Registration(database *sql.DB, userArguments UserArguments) error {
	// Выполняем запрос и возвращаем результат
	_, err := database.Exec(registrationNewUser, userArguments.Login, userArguments.Password)
	return err
}

/*
Authorization - запрос на авторизацию пользователя
*/
func Authorization(database *sql.DB, userArguments UserArguments) (UserData, error) {
	// Выполняем запрос
	rows, err := database.Query(authorizationUser, userArguments.Login, userArguments.Password)
	defer rows.Close()

	if err != nil {
		return UserData{}, err
	}

	// Читаем полученную информацию и возвращаем её
	userData := UserData{
		Nickname: userArguments.Login,
	}
	for rows.Next() {
		rows.Scan(&userData.AccessToken, &userData.Status)
	}

	return userData, nil
}

/*
CheckAccess - запрос на проверку актуальности токена
*/
func CheckAccess(database *sql.DB, accessToken string) (UserData, error) {
	// Выполняем запрос
	rows, err := database.Query(testUserSessionKey, accessToken)
	defer rows.Close()

	if err != nil {
		return UserData{}, err
	}

	// Читаем полученную информацию и возвращаем её
	userData := UserData{
		AccessToken: accessToken,
	}

	for rows.Next() {
		rows.Scan(&userData.Nickname, &userData.Status)
	}

	return userData, nil
}

/*
Logout - закрытие сессии пользователя
*/
func Logout(database *sql.DB, accsess_token string) (bool, error) {
	// Выполняем запрос и возвращаем результат
	_, err := database.Exec(logoutUser, accsess_token)
	return err == nil, err
}
