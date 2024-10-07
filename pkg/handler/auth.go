package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/kingxl111/RESTapiService"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	// в input хотим распарсить тело JSON который получили от пользователя, поэтому передаем его по ссылке
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}
	
	// Создаем пользователя
	id, err := h.services.Authorization.CreateUser(input)
	// Если не удалось создать нового пользователя, надо сообщить ему, что ошибка на стороне сервера
	// То есть вернуть ошибку 500 - internalError
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	// Если всё хорошо, то будет статус 200 - StatusOK
	c.JSON(http.StatusOK, map[string]interface{}{
		"id" : id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	// в input хотим распарсить тело JSON который получили от пользователя, поэтому передаем его по ссылке
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}
	
	// Создаем пользователя
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	// Если не удалось создать нового пользователя, надо сообщить ему, что ошибка на стороне сервера
	// То есть вернуть ошибку 500 - internalError
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	// Если всё хорошо, то будет статус 200 - StatusOK
	c.JSON(http.StatusOK, map[string]interface{}{
		"token" : token,
	})
}