package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func initRouter() {

	port := getEnvVar("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/users/{username}", getUser).Methods("GET")
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/delete/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/ask-a-question", askAQuestion).Methods("POST")
	router.HandleFunc("/users/answer", answerQuestion).Methods("POST")
	router.HandleFunc("/questions/get/{question_id}", getQuestion).Methods("GET")

	//AUTH
	router.HandleFunc("/auth/register", registerUser).Methods("POST")
	router.HandleFunc("/auth/login", loginUser).Methods("POST")
	router.HandleFunc("/auth/userExists/{username}", userExists).Methods("GET")

	http.ListenAndServe((":" + port),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET, POST, PUT, DELETE, OPTIONS"}),
		)(router))
}

func main() {
	initDatabase()
	initRouter()
}
