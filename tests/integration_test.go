package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"projekt/internal/handlers"
	"projekt/internal/repository"
	"projekt/internal/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	repo := repository.New()
	bLayer := services.New(repo)
	handler := handlers.New(bLayer)

	router := gin.Default()
	router.GET("/users", handler.GetUsersHandler)
	router.GET("/users/:id", handler.GetUserHandler)
	router.POST("/users", handler.AddUserHandler)
	router.PATCH("/users/:id", handler.UpdateUserHandler)
	router.DELETE("/users/:id", handler.DeleteUserHandler)

	return router
}

func TestCreateUser(t *testing.T) {
	router := setupRouter()
	userPayload := `{"firstName": "Dawid", "lastName": "Markiewicz", "birthYear": 2009, "group": "user"}`

	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(userPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserByID(t *testing.T) {
	router := setupRouter()

	userPayload := `{"firstName": "Dawid", "lastName": "Markiewicz", "birthYear": 2009, "group": "user"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(userPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"firstName":"Dawid"`)
}

func TestUpdateUser(t *testing.T) {
	router := setupRouter()

	userPayload := `{"firstName": "Dawid", "lastName": "Markiewicz", "birthYear": 2009, "group": "user"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(userPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	updatePayload := `{"firstName": "aaaa", "lastName": "dadw", "birthYear": 2009, "group": "user"}`
	req, _ = http.NewRequest("PATCH", "/users/1", bytes.NewBufferString(updatePayload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	router := setupRouter()

	userPayload := `{"firstName": "Dawid", "lastName": "Markiewicz", "birthYear": 2009, "group": "user"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(userPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	req, _ = http.NewRequest("DELETE", "/users/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	req, _ = http.NewRequest("GET", "/users/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
