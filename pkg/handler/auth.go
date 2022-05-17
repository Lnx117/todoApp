package handler

import (
	"net/http"
	"todo"

	"github.com/gin-gonic/gin"
)

//Тут мы обрабатываем запрос из handler.go, распарсиваем то что нам приходит и отсылаем на уровень бизнес логики
func (h *Handler) signUp(c *gin.Context) {
	//переменная со структурой из файла, в ней храним то что посылает нам пользователь в запросе
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Тут services это уже готовый объект реализующий интерфейс который нам надо (реализация интерфейса произошла еще при создании
	//объекта в main). В Authorization лежит объект из конструктора файла  auth.go конструктор NewAuthService
	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//Структура запроса. Мы получаем логин и пароль от пользователя
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	//переменная со структурой того что хотим получить
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Тут services это уже готовый объект реализующий интерфейс который нам надо (реализация интерфейса произошла еще при создании
	//объекта в main). В Authorization лежит объект из конструктора файла  auth.go конструктор NewAuthService
	//Генерируем токен
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//возвращаем токен
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
