package main

import (
	"log"

	"github.com/joho/godotenv"
	"gitlab.liu.se/adaab301/tddd27_2022_project/backend/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	api.Start()
}
