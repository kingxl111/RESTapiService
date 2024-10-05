package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kingxl111/RESTapiService/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Метод инициализирующий endPoint'ы 
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Группируем по маршрутам
	auth := router.Group("/auth") 
	{
		// Задаем http-метод, который будет использоваться при обращении к заданному end-point'у
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}

	// Непосредственно маршрутизация внутри нашего основного функционала
	api := router.Group("/api") 
	{
		// Создаем еще одну группу для работы со списками todo
		lists := api.Group("/lists")
		{
			// задаем "relativePath"
			lists.POST("/", h.createList)					// Создание нового списка 
			lists.GET("/", h.getAllLists)					// Получние всех списков

			// Двоеточие означает, что там может быть любое значение параметра id 
			lists.GET("/:id", h.getListById) 				// Получение списка по id из URL
			lists.PUT("/:id", h.updateList)					// обновление списка по id из URLа
			lists.DELETE("/:id", h.deleteList)				// удаление списка по id из URL

			items := lists.Group("/items")
			{
				items.POST("/", h.createItem)				// Создание нового элемента
				items.GET("/", h.getAllItems)				// Получние всех элементов списка
				
				items.GET("/:item_id", h.getItemById)		// Получние элемента по id
				items.PUT("/:item_id", h.updateItem)		// Обновление элемениа по id
				items.DELETE("/:item_id", h.deleteItem)		// Удалние элемента по id 
			}

		}
	}
	return router

}