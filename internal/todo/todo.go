// internal/todo/todo.go

package todo

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}


type TodoService struct {
	Todos []Todo
}

func NewTodoService() *TodoService {
	return &TodoService{
		Todos: []Todo{},
	}
}

func (ts *TodoService) GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ts.Todos)
}

func (ts *TodoService) AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	newTodo.ID = len(ts.Todos) + 1
	ts.Todos = append(ts.Todos, newTodo)
	w.WriteHeader(http.StatusCreated)
}

func (ts *TodoService) RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoID struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&todoID); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, todo := range ts.Todos {
		if todo.ID == todoID.ID {
			ts.Todos = append(ts.Todos[:i], ts.Todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "To-Do item not found", http.StatusNotFound)
}


func (ts *TodoService) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, todo := range ts.Todos {
		if todo.ID == updatedTodo.ID {
			ts.Todos[i] = updatedTodo
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "To-Do item not found", http.StatusNotFound)
}

func (ts *TodoService) CompleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todoID struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&todoID); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i, todo := range ts.Todos {
		if todo.ID == todoID.ID {
			ts.Todos[i].Completed = true
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "To-Do item not found", http.StatusNotFound)
}
