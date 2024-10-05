Запускаем докер-образ с postgreSQL
sudo docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres


Создаем файлы миграции базы данных
migrate create -ext sql -dir ./schema -seq init

