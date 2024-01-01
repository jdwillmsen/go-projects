package main

import (
	"encoding/json"
	"fmt"
	"github.com/jdwillmsen/todo-api/todolib"
	"net/http"
	"strconv"
)

var todoStore *todolib.TodoStore

func init() {
	todoStore = todolib.NewTodoStore()
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo todolib.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todoID := todoStore.Create(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, "{\"id\": %d}", todoID)
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
	}

	todo, exists := todoStore.Read(id)
	if !exists {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(todo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	var updatedTodo todolib.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !todoStore.Update(id, updatedTodo) {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	if !todoStore.Delete(id) {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func listTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := todoStore.List()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(todos)
}

func main() {
	port := "9406"

	http.HandleFunc("/api/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTodoHandler(w, r)
		case http.MethodGet:
			getTodoHandler(w, r)
		case http.MethodPut:
			updateTodoHandler(w, r)
		case http.MethodDelete:
			deleteTodoHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			listTodosHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server listening on port " + port)
	_ = http.ListenAndServe(":"+port, nil)
}
