package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/one-d-plate/go-boilerplate.git/cmd"
	"github.com/one-d-plate/go-boilerplate.git/src/pkg"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pkg.SignalInit()
	pkg.InitLogger()
	cmd.Execute()
}
