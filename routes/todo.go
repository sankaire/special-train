package routes

import (
	"github.com/gorilla/mux"
	"todo-api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", controllers.Home).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/new/todo", controllers.CreateTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/get/todos", controllers.GetAllTask).Methods("GET", "OPTIONS")
	return router
}
