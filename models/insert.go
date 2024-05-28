package models

import (
	"github.com/CamposLeo95/projeto_go.git/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	// ? defer somente e rodado depois que a funcao concluir
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	//! Pesquisar o que essa linha faz exatamente
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
