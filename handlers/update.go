package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CamposLeo95/projeto_go.git/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {

	// Pegar o Id do parâmetro da URL
	//? strconv.Atoi: Converte a string para um int
	//? chi: usado para puxar um parâmetro da URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Inserir a variável todo
	var todo models.Todo

	// Decodificar o JSON do body
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Fazer o Update e checar se há erro
	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao fazer update do todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Verificar se mais de uma linha foi atualizada
	if rows > 1 {
		log.Printf("Erro: foram atualizados %d registros", rows)
	}

	// Caso passe por tudo, montar a resposta com map
	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo atualizado com sucesso! ID: %d", id),
	}

	// Estruturar a resposta com um header JSON e criar um JSON com a resposta
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
