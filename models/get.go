package models

import "github.com/CamposLeo95/projeto_go.git/db"

func Get(id int64) (todo Todo, err error) {

	//primeiro passo abrir connexao com o DB
	conn, err := db.OpenConnection()
	// segundo passo verificar se possui erro
	if err != nil {
		return
	}
	// terceito : fechar conexao com o banco usando defer para executar somente no final
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
