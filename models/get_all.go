package models

import "github.com/CamposLeo95/projeto_go.git/db"

func GetAll() (todos []Todo, err error) {
	// TODO[X] Abrir conexao com o banco
	conn, err := db.OpenConnection()

	// TODO[X] Checar se ha erro
	if err != nil {
		return
	}

	// TODO[X] Defer (fechar conexao com o db)
	defer conn.Close()

	// TODO[X] fazer a query para o banco
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	// TODO[X] fazer um for quando for varios retornos
	for rows.Next() {
		var todo Todo

		// TODO[X] fazer o Scan() e ver err
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		// TODO[X] checar erro e se der continuar
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	// TODO[X] Return
	return

}
