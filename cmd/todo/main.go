// cmd/todo/main.go

package main

import (
	"log"
	"net/http"
	"todo-microservice/internal/todo"
)

func main() {
	todoService := todo.NewTodoService()
	http.HandleFunc("/todos", todoService.GetTodosHandler)
	http.HandleFunc("/todos/add", todoService.AddTodoHandler)
	http.HandleFunc("/todos/remove", todoService.RemoveTodoHandler)
	http.HandleFunc("/todos/update", todoService.UpdateTodoHandler)
	http.HandleFunc("/todos/complete", todoService.CompleteTodoHandler)
	http.Handle("/", http.FileServer(http.Dir("web")))
	log.Println("Starting server on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
