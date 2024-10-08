package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/kingxl111/RESTapiService"
)

// В этом handler'е будет обработка списков todo
// Все методы будут отвечать контракту основного handler
/*
	lists.POST("/")			// Создание нового списка
	lists.GET("/")			// Получние всех списков

	// Двоеточие означает, что там может быть любое значение параметра id
	lists.GET("/:id") 		// Получение списка по id из URL
	lists.PUT("/:id")		// обновление списка по id из URLа
	lists.DELETE("/:id")	// удаление списка по id из URL
*/

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	// call service 
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// call service 
	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})

}


func (h *Handler) getListById(c *gin.Context) {

}


func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}