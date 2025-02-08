package main

import (
	"projekt/internal/repository"
	"projekt/internal/services"
	"projekt/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.New()
	bLayer := services.New(repo)
	handler := handlers.New(bLayer)

	router := gin.New()

	router.GET("users", handler.GetUsersHandler)
	router.GET("users/:id", handler.GetUserHandler)

	router.POST("users", handler.AddUserHandler)
	router.PATCH("users/:id", handler.UpdateUserHandler)
	router.DELETE("users/:id", handler.DeleteUserHandler)
}
