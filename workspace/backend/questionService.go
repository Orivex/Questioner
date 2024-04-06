package main

import "net/http"

func questionExists(w http.ResponseWriter, question Question) bool {
	err := DB.First(&question, "id = ?", question.ID).Error
	if err != nil {
		http.Error(w, "Question doesn't exist", 400)
		return false
	}

	return true
}
