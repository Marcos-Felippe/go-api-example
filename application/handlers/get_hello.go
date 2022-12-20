package handlers

import (
	"encoding/json"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {

	res := "Hello from API"

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
