package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVar(variable string) string {
	godotenv.Load(".env")

	res := os.Getenv(variable)
	if res == "" {
		log.Fatal(variable, " not bound")
		return ""
	} else {
		return res
	}
}
