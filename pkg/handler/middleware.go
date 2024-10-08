// По сути прослойка, чтобы спарсить JWT токен и авторизовать пользователя
package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

// Это разумеется метод обработчика
// Здесь нужно получать значение из хедера авторизации, валидировать его, парсить токен 
// и записывать пользователя в контекст
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return 
	}
	
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return 
	}
	

	// При ошибках возвращаем статус-код 401, что означает, что пользователь не авторизирован
	
	// Парсим токен
	// fmt.Println(headerParts[1])
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return 
	}

	c.Set(userCtx, userId)

	// Получается, если операция успешная, то мы записываем id в контекст
	// Это делается для того, чтобы иметь доступ к id пользователя в последующих обработчиках
	// которые вызываются после данной прослойки
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}