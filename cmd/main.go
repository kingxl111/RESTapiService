package main

import (
	"log"

	todo "github.com/kingxl111/RESTapiService"
	"github.com/kingxl111/RESTapiService/pkg/handler"
	"github.com/kingxl111/RESTapiService/pkg/repository"
	"github.com/kingxl111/RESTapiService/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	// "database/sql"
	// "github.com/golang-migrate/migrate/v4"
    // "github.com/golang-migrate/migrate/v4/database/postgres"


    _ "github.com/golang-migrate/migrate/v4/source/file"
    _ "github.com/lib/pq"
)

/*
func migrat() {
	dbURL := "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable"
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Could not open database: %v", err)
    }
	
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        log.Fatalf("Could not create driver: %v", err)
    }
	
    m, err := migrate.NewWithDatabaseInstance(
        "file://schema",
        "postgres", driver)
		if err != nil {
			log.Fatalf("Could not create migrate instance: %v", err)
		}
		
		// Применяем миграции
		if err := m.Up(); err != nil {
			log.Fatalf("Could not apply migrations: %v", err)
		}
		
		log.Println("Migrations applied successfully")
}	

// sudo docker exec -it todo-db psql -U postgres -d postgres

 Schema |        Name        |   Type   |  Owner   
--------+--------------------+----------+----------
 public | lists_items        | table    | postgres
 public | lists_items_id_seq | sequence | postgres
 public | schema_migrations  | table    | postgres
 public | todo_items         | table    | postgres
 public | todo_items_id_seq  | sequence | postgres
 public | todo_lists         | table    | postgres
 public | todo_lists_id_seq  | sequence | postgres
 public | user_lists         | table    | postgres
 public | user_lists_id_seq  | sequence | postgres
 public | users              | table    | postgres
 public | users_id_seq       | sequence | postgres
*/


func main() {

	// migrat()

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
