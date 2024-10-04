package main

import (
	"log"
	"github.com/kingxl111/RESTapiService"
	"github.com/kingxl111/RESTapiService/pkg/handler"
)

func main() {

	// Создаем handlers. Чтобы всё работало, нужно наличие хотя бы одного handler
	// класс Handler находится в файле handler.go
	handlers := new(handler.Handler)

	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("error appeared while running http server: %s", err.Error())
	}

}