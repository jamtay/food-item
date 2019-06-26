package response_handler

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if msg == "not found"{
		code = http.StatusNotFound
	}
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
