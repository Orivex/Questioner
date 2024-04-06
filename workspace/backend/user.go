package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Questions []Question `gorm:"foreignKey:UserID"`
}

func answerQuestion(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)

	type Request struct {
		Answer      string `json:"answer"`
		Question_ID int    `json:"question_id"`
	}

	var data Request
	json.NewDecoder(r.Body).Decode(&data)

	var question Question
	DB.First(&question, "id = ?", data.Question_ID)

	question.Answer = data.Answer
	DB.Save(&question)
	respondWithJson(w, "Question answered successfully")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)
	params := mux.Vars(r)
	fmt.Printf("Received parameters: %+v\n", params)
	var user User
	DB.First(&user, "username = ?", params["username"])
	DB.Preload("Questions").First(&user)
	respondWithJson(w, &user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)
	var users []User
	DB.Find(&users)
	respondWithJson(w, &users)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Print("Users username:", user.Username)
	user.Password = randomPassword()
	DB.Create(&user)
	respondWithJson(w, &user)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	var userFromDb User
	DB.First(&userFromDb, "username = ?", user.Username)
	if !checkPassword(userFromDb.Password, user.Password) {
		http.Error(w, "False password", 400)
	}
	respondWithJson(w, "Logged in")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	respondWithJson(w, json.NewEncoder(w).Encode("User deleted"))
}
