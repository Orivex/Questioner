package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Println("Something went wrong")
	}

	w.Write(data)
}

func setContentTypeToJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
