package controllers

import (
	"fmt"
	"net/http"

	"./database"
	"./filesystem"
)

/*
RegistrationHandler - обрабатывает запрос на регистрацию пользователя
*/
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Создаём структуру с пользовательским логином и паролем
	user := database.UserArguments{}
	// Получаем информацию о логине и пароле в структуру
	err := jsonFromBody(r, &user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Делаем запрос к БД
	err = database.Registration(gDatabase, user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		return
	}

	// Если всё ок - создаём для него папку
	filesystem.CreateFolder(filesystem.RootPath, user.Login)
}

/*
AuthorizationHandler - обрабатывает запрос на авторизацию пользователя
*/
func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	// Создаём структуру с пользовательским логином и паролем
	user := database.UserArguments{}
	// Получаем информацию о логине и пароле в структуру
	err := jsonFromBody(r, &user)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		w.Header().Add("Error", err.Error())
		return
	}

	// Выполняем авторизацию и получаем данные пользователя или ошибку
	userData, err := database.Authorization(gDatabase, user)
	if err != nil {
		// Если ошибка - отправляем ошибку
		makeErrorHeader(w, http.StatusBadRequest)
	} else {
		// Иначе отправляем JSON с пользовательскими данными
		makeJsonHeader(w, "POST", userData)
	}
}

/*
TestTokenHandler - обрабатывает запрос на проверку валидности токена
*/
func TestTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Создаём структуру для токена пользователя
	token := database.Token{}
	// Получаем токен из запроса
	err := jsonFromBody(r, &token)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
		w.Header().Add("Error", err.Error())
		return
	}

	// Проверяем валидность токена
	userData, err := checkAccess(token.Token)
	if err != nil {
		makeErrorHeader(w, http.StatusBadRequest)
	} else {
		// Возвращаем информацию о пользователе, если всё ок
		makeJsonHeader(w, "POST", userData)
	}
}

/*
LogoutHandler - обрабатывает запрос на завершение сессии
*/
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем токен пользователя из запроса
	token := database.Token{}
	err := jsonFromBody(r, &token)
	if err != nil {
		fmt.Println(err)
		makeErrorHeader(w, http.StatusBadRequest)
		w.Header().Add("Error", err.Error())
		return
	}

	// Делаем запрос к БД
	database.Logout(gDatabase, token.Token)
}
