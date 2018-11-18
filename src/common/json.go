package common

import (
	"net/http"
	"encoding/json"
	"strings"
	"errors"
	"bytes"
	"github.com/amstee/easy-cut/src/logger"
	"github.com/amstee/easy-cut/src/types"
)

func JsonToReader(data interface{}) (*bytes.Buffer, error) {
	jsonData, err := json.Marshal(data); if err != nil {
		logger.Error.Println(err)
		return nil, err
	}
	return bytes.NewBuffer(jsonData), nil
}

func GetBearer(r *http.Request) (string, error) {
	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authHeaderParts) < 2 {
		return "", errors.New("authorization token not found")
	}
	if authHeaderParts[0] == "Bearer" {
		tokenString := authHeaderParts[1]
		return tokenString,  nil
	}
	return "", errors.New("invalided identifier found")
}

func ResponseError(message string, err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		logger.Error.Println(message + " : " + err.Error())
		ResponseJSON(types.HttpMessage{Message: message + " : " + err.Error(), Success: false}, w, statusCode)
	} else {
		logger.Error.Println(message + " : nil error")
		ResponseJSON(types.HttpMessage{Message: message, Success: false}, w, statusCode)
	}
}

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