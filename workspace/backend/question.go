package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID          uint
	Questioner  string `json:"questioner"`
	TheQuestion string `json:"thequestion"`
	Answer      string `json:"answer"`
	UserID      uint   `json:"user_id"`
}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)
	params := mux.Vars(r)
	var question Question
	DB.First(&question, "id = ?", params["question_id"])
	respondWithJson(w, &question)
}

func askAQuestion(w http.ResponseWriter, r *http.Request) {
	setContentTypeToJson(w)

	type Request struct {
		Questioner  string `json:"questioner"`
		Thequestion string `json:"thequestion"`
		Username    string `json:"username"`
	}
	var data Request
	json.NewDecoder(r.Body).Decode(&data)
	log.Println(data)

	var user User
	if err := DB.First(&user, "username = ?", data.Username).Error; err != nil {
		http.Error(w, "User doesn't exist", 400)
		return
	}

	var question Question
	question.TheQuestion = data.Thequestion
	question.Questioner = data.Questioner
	question.UserID = user.ID
	DB.Create(&question)

	user.Questions = append(user.Questions, question)
	DB.Save(&user)
	respondWithJson(w, "Question asked...")
}
