package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/projetosgo/exemploapi/application/database"
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/application/usecase"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	db, err := database.Database()
	if err != nil {
		log.Printf("Erro ao se conectar com o banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Printf("Erro ao pegar id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var userInput usecase.UserInputDTO
	err = json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		log.Printf("Erro ao fazer decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	repository := repository.NewUserRepository(db)
	uc := usecase.NewUpdateUserUseCase(repository)

	outputDTO, err := uc.Execute(id, userInput)
	if err != nil {
		log.Printf("Erro ao salvar usuario no banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputDTO)

}
