package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/projetosgo/exemploapi/application/database"
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/application/usecase"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	db, err := database.Database()
	if err != nil {
		log.Printf("Erro ao se conectar com o banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var userInput usecase.UserInputDTO
	err = json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		log.Printf("Erro ao fazer decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	repository := repository.NewUserRepository(db)
	uc := usecase.NewCreateUserUseCase(repository)

	outputDTO, err := uc.Execute(userInput)
	if err != nil {
		log.Printf("Erro ao salvar usuario no banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputDTO)

}
