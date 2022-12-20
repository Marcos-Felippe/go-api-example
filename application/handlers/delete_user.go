package handlers

import (
	"log"
	"net/http"

	"github.com/projetosgo/exemploapi/application/database"
	"github.com/projetosgo/exemploapi/application/database/repository"
	"github.com/projetosgo/exemploapi/application/usecase"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

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

	repository := repository.NewUserRepository(db)
	uc := usecase.NewDeleteUserUseCase(repository)

	err = uc.Execute(id)
	if err != nil {
		log.Printf("Erro ao deletar usuario do banco: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	return

}
