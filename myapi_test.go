package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUsers(t *testing.T) {
	// Create a new router.
	r := gin.Default()

	// Define the routes.
	r.GET("/users", getUsers)

	// Create a new request.
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder.
	w := httptest.NewRecorder()

	// Serve the request.
	r.ServeHTTP(w, req)

	// Check the status code.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body.
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		t.Fatal(err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	if users[0].Name != "John Doe" {
		t.Errorf("Expected user name to be \"John Doe\", got \"%s\"", users[0].Name)
	}

	if users[0].Email != "johndoe@example.com" {
		t.Errorf("Expected user email to be \"johndoe@example.com\", got \"%s\"", users[0].Email)
	}

	if users[1].Name != "Jane Doe" {
		t.Errorf("Expected user name to be \"Jane Doe\", got \"%s\"", users[1].Name)
	}

	if users[1].Email != "janedoe@example.com" {
		t.Errorf("Expected user email to be \"janedoe@example.com\", got \"%s\"", users[1].Email)
	}
}

func TestGetUser(t *testing.T) {
	// Create a new router.
	r := gin.Default()

	// Define the routes.
	r.GET("/users/:id", getUser)

	// Create a new request.
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder.
	w := httptest.NewRecorder()

	// Serve the request.
	r.ServeHTTP(w, req)

	// Check the status code.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body.
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	if user.Name != "John Doe" {
		t.Errorf("Expected user name to be \"John Doe\", got \"%s\"", user.Name)
	}

	if user.Email != "johndoe@example.com" {
		t.Errorf("Expected user email to be \"johndoe@example.com\", got \"%s\"", user.Email)
	}
}

func TestCreateUser(t *testing.T) {
	// Create a new router.
	r := gin.Default()

	// Define the routes.
	r.POST("/users", createUser)

	// Create a new request.
	req, err := http.NewRequest("POST", "/users", bytes.NewReader([]byte(`{"name": "John Doe", "email": "johndoe@example.com"}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder.
	w := httptest.NewRecorder()

	// Serve the request.
	r.ServeHTTP(w, req)

	// Check the status code.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body.
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	if user.Name != "John Doe" {
		t.Errorf("Expected user name to be \"John Doe\", got \"%s\"", user.Name)
	}

	if user.Email != "johndoe@example.com" {
		t.Errorf("Expected user email to be \"johndoe@example.com\", got \"%s\"", user.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	// Create a new router.
	r := gin.Default()

	// Define the routes.
	r.PUT("/users/:id", updateUser)

	// Create a new request.
	req, err := http.NewRequest("PUT", "/users/1", bytes.NewReader([]byte(`{"name": "Jane Doe", "email": "janedoe@example.com"}`)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder.
	w := httptest.NewRecorder()

	// Serve the request.
	r.ServeHTTP(w, req)

	// Check the status code.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body.
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal(err)
	}

	if user.Name != "Jane Doe" {
		t.Errorf("Expected user name to be \"Jane Doe\", got \"%s\"", user.Name)
	}

	if user.Email != "janedoe@example.com" {
		t.Errorf("Expected user email to be \"janedoe@example.com\", got \"%s\"", user.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	// Create a new router.
	r := gin.Default()

	// Define the routes.
	r.DELETE("/users/:id", deleteUser)

	// Create a new request.
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder.
	w := httptest.NewRecorder()

	// Serve the request.
	r.ServeHTTP(w, req)

	// Check the status code.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body.
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	var message Message
	if err := json.Unmarshal(body, &message); err != nil {
		t.Fatal(err)
	}

	if message.Message != "User deleted" {
		t.Errorf("Expected message to be \"User deleted\", got \"%s\"", message.Message)
	}
}
