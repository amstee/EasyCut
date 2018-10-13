package utils

import (
	"net/http"
	"encoding/json"
)

func ResponseJSON(data interface{}, w http.ResponseWriter, statusCode int) {
	jsonResponse, err := json.Marshal(data); if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func DecodeJSON(dest interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(dest); if err != nil {
		return err
	}
	return nil
}