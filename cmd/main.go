package main

import (
	"log"

	"github.com/kingxl111/RESTapiService"
	"github.com/kingxl111/RESTapiService/pkg/handler"
	"github.com/kingxl111/RESTapiService/pkg/repository"
	"github.com/kingxl111/RESTapiService/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	// Инициализируем конфиги
	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)

	// Создаем handlers. Чтобы всё работало, нужно наличие хотя бы одного handler
	// класс Handler находится в файле handler.go
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error appeared while running http server: %s", err.Error())
	}
	
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}