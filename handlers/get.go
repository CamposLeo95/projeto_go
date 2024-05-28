package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CamposLeo95/projeto_go.git/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar Todo")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(todo)
}
