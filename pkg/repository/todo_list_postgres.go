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


func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	
	// INNER JOIN помогает выбрать только те записи, которые имеют одинаковые значение в обеих таблицах
	// То есть мы для каждого пользователя хотим выделить все его списки
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListTable, usersListsTable)
	
	// Select работает аналогично методу Get, но применяется при выборке больше одного элемента и записи в slice
	err := r.db.Select(&lists, query, userId)
	
	return lists, err
}

func (r *TodoListPostgres) GetList(userId, listId int) (todo.TodoList, error) {

	var lst todo.TodoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl 
						  INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 
						  AND ul.list_id = $2`, todoListTable, usersListsTable)

	err := r.db.Get(&lst, query, userId, listId)
	
	return lst, err
}