$
sudo nano /etc/resolv.conf 
nameserver 8.8.8.8
nameserver 8.8.4.4
$

Запуск контейнера с psql
docker run --name todo-db -e POSTGRES_PASSWORD=qwerty -d -p 5432:5432 postgres


Создаем файлы миграции базы данных
migrate create -ext sql -dir ./schema -seq init

Удобная библиотека для работы с бд на основе database/sql
go get github.com/jmoiron/sqlx


Для работы с переменными окружения(сейчас конкретно для пароля)
go get github.com/joho/godotenv

Добавим библиотеку с логгером 
go get github.com/sirupsen/logrus
Для логгера лучше использовать JSON формат


Важный момент: исходную миграцию с созданием всех таблиц базы данных
нужно сделать единожды после создание контейнера с postgreSQL
sudo docker start <container_name>
sudo docker exec -it <container_id> /bin/bash
psql -U <database_user>

Библиотека для работы с токенами JWT
go get -u github.com/golang-jwt/jwt/v5