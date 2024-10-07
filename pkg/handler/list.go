package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	id, _ := c.Get(userCtx) // const userCtx здесь задается в middleware(назвать его можно как угодно)
 	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}


func (h *Handler) getAllLists(c *gin.Context) {

}


func (h *Handler) getListById(c *gin.Context) {

}


func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}