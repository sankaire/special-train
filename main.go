package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-api/database"
	"todo-api/routes"
)

func main() {
	r := routes.Router()
	database.Init()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
