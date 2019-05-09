package controllers

import (
	"encoding/json"
	"net/http"
)

func makeJsonHeader(w http.ResponseWriter, method string, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", method)
	w.Header().Set("Access-Control-Allow-Headers", "application/json")
	err := json.NewEncoder(w).Encode(obj)

	return err
}

func makeErrorHeader(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
