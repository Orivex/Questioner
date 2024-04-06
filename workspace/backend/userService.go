package main

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func userExists(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	if err := DB.First(&user, "username = ?", params["username"]).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithJson(w, false)
			return
		}

	}

	respondWithJson(w, true)
}

func checkPassword(realPassword string, password string) bool {
	if realPassword != password {
		return false
	}

	return true
}

func randomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!$^()_;"
	password := make([]byte, 10)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
