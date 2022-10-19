package controllers

import (
	"encoding/json"
	"net/http"
	"todo-api/database"
	"todo-api/models"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.TodoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	database.InsertOne(task)
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		return
	}
}
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := database.Getalltasks()
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	var res = map[string]interface{}{
		"message": "Welcome to todo api",
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
