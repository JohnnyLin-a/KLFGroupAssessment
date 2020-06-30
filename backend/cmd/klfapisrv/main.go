package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	fmt.Println("POSTGRES_USER", os.Getenv("POSTGRES_USER"))
	fmt.Println("POSTGRES_PASSWORD", os.Getenv("POSTGRES_PASSWORD"))
	fmt.Println("POSTGRES_DB", os.Getenv("POSTGRES_DB"))

	for true {
	}
}
