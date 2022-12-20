package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/projetosgo/exemploapi/application/database"
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/application/usecase"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	db, err := database.Database()
	if err != nil {
		log.Printf("Erro ao se conectar com o banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	uc := usecase.NewGetAllUseCase(repository)

	users, err := uc.Execute()
	if err != nil {
		log.Printf("Erro ao pegar usuarios no banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
