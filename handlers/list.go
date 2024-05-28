package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CamposLeo95/projeto_go.git/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Erro ao carregar todos!")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(todos)
}
