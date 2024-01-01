package main

import (
	"bytes"
	"encoding/json"
	"github.com/jdwillmsen/todo-api/todolib"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodoHandler(t *testing.T) {
	todoStore = todolib.NewTodoStore()

	todo := todolib.Todo{
		Title:     "Test Todo",
		Completed: false,
	}
	todoJson, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}
}

func TestGetTodoHandler(t *testing.T) {
	todoStore = todolib.NewTodoStore()
	expectedTodo := todolib.Todo{
		ID:        1,
		Title:     "Test Todo",
		Completed: false,
	}
	todoJson, _ := json.Marshal(expectedTodo)
	req, err := http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/api/todo?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	http.HandlerFunc(getTodoHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var returnedTodo todolib.Todo
	err = json.NewDecoder(rr.Body).Decode(&returnedTodo)
	if err != nil {
		t.Error(err.Error())
	}

	if returnedTodo != expectedTodo {
		t.Errorf("Expected todo %v, got %v", expectedTodo, returnedTodo)
	}
}

func TestUpdateTodoHandler(t *testing.T) {
	todoStore = todolib.NewTodoStore()
	todo := todolib.Todo{
		ID:        1,
		Title:     "Test Todo",
		Completed: false,
	}
	todoJson, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	updatedTodo := todolib.Todo{
		ID:        1,
		Title:     "Updated Test Todo",
		Completed: true,
	}
	updatedTodoJson, _ := json.Marshal(updatedTodo)
	req, err = http.NewRequest("PUT", "/api/todo?id=1", bytes.NewReader(updatedTodoJson))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	http.HandlerFunc(updateTodoHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rr.Code)
	}
}

func TestDeleteTodoHandler(t *testing.T) {
	todoStore = todolib.NewTodoStore()
	todo := todolib.Todo{
		ID:        1,
		Title:     "Test Todo",
		Completed: false,
	}
	todoJson, _ := json.Marshal(todo)
	req, err := http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	req, err = http.NewRequest("DELETE", "/api/todo?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	http.HandlerFunc(deleteTodoHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rr.Code)
	}
}

func TestListTodosHandler(t *testing.T) {
	todoStore = todolib.NewTodoStore()
	expectedTodos := []todolib.Todo{
		{
			ID:        1,
			Title:     "Test Todo 1",
			Completed: false,
		},
		{
			ID:        2,
			Title:     "Test Todo 2",
			Completed: false,
		},
		{
			ID:        3,
			Title:     "Test Todo 3",
			Completed: false,
		},
	}
	todoJson1, _ := json.Marshal(expectedTodos[0])
	req, err := http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson1))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	todoJson2, _ := json.Marshal(expectedTodos[1])
	req, err = http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson2))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	todoJson3, _ := json.Marshal(expectedTodos[2])
	req, err = http.NewRequest("POST", "/api/todo", bytes.NewReader(todoJson3))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	http.HandlerFunc(createTodoHandler).ServeHTTP(rr, req)

	req, err = http.NewRequest("GET", "/api/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	http.HandlerFunc(listTodosHandler).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var returnedTodos []todolib.Todo
	err = json.NewDecoder(rr.Body).Decode(&returnedTodos)
	if err != nil {
		t.Error(err.Error())
	}

	if len(returnedTodos) != len(expectedTodos) {
		t.Errorf("Expected todo %v, got %v", expectedTodos, returnedTodos)
	}
}
