package utils

import(
	"net/http"
	"encoding/json"
)

func WriteJSON(w http.ResponseWriter,data any){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}