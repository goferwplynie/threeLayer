package handlers

import (
	"net/http"
	"projekt/internal/models"
	"projekt/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	BLayer *services.BuisnessLayer
}


func (h *Handler) AddUserHandler(ctx *gin.Context){
	var request models.Request	

	err := ctx.ShouldBindJSON(&request)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = h.BLayer.AddUser(request)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) GetUsersHandler(ctx *gin.Context){
	users := h.BLayer.GetAllUsers

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *Handler) GetUserHandler(ctx *gin.Context){
	
}
