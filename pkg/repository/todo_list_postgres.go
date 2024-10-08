package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/kingxl111/RESTapiService"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

// Будем реализовывать транзакцию.
// Транзакция - это последовательность нескольких операций
// В данном случае нужно добавить новый список в таблицы todo_lists и users_lists
// Это получается две операции, но одна транзакция
func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	// Создаем новую запись в таблице todo_lists 
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	// Создаем новую запись в таблице users_lists
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	// Для простого выполнения запроса без чтения дополнительной информации воспользуемся методом Exec
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		// Rollback откатывает все изменения базы данных до начала выполнения транзакции
		tx.Rollback()
		return 0, err
	}

	// вызываем Commit, который применяет изменения к базе данных, и заканчиваем транзакцию
	return id, tx.Commit()
}
