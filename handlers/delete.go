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

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Error ao deletar todo: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if rows > 1 {
		fmt.Sprintf("Erro: Foram deletadas %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo deletado com sucesso"),
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(resp)
}
