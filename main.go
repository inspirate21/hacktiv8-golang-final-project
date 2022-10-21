package main

import (
	"log"

	"hacktiv8-golang-final-project/server"

	_ "github.com/joho/godotenv/autoload"
)

// @title Final Project Hacktiv8 Golang
// @version 1.0
// @description This is API for final project hacktiv8 golang

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
