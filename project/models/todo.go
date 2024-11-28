package models

import (
	"log"
	"todo/database"
)

type Todo struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Completed   bool   `db:"completed" json:"completed"`
}

func GetAllTodos(page int64, pageSize int64) (*[]Todo, error) {
	var todos []Todo
	offset := (page - 1) * pageSize
	query := "SELECT * FROM todo LIMIT $1 OFFSET $2"
	err := database.DB.Select(&todos, query, pageSize, offset)

	if err != nil {
		return nil, err
	}
	return &todos, nil
}

func AddTodos(t Todo) error {
	_, err := database.DB.Exec(`INSERT INTO todo(title, description, completed) VALUES ($1, $2, $3)`, t.Title, t.Description, t.Completed)

	if err != nil {
		log.Println("Error inserting todo:", err)
		return err
	}
	return nil
}

func GetTodoById(id int64) (*Todo, error) {
	query := "SELECT * FROM todo WHERE id = $1"
	row := database.DB.QueryRow(query, id)

	var todo Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (todo Todo) DeleteTodoById() error {
	query := "DELETE FROM todo WHERE id = $1"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(todo.ID)
	return nil
}

func (todo Todo) UpdateTodo() error {
	query := "UPDATE todo SET title=$1, description=$2, completed=$3 WHERE id=$4"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, todo.Completed, todo.ID)

	return err
}

func DeleteAll() error {
	query := "TRUNCATE todo RESTART IDENTITY"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
