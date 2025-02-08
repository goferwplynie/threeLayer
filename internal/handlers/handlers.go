package handlers

import (
	"net/http"
	"projekt/internal/models"
	"projekt/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	BLayer *services.BuisnessLayer
}

func New(bLayer *services.BuisnessLayer) *Handler{ 
	return &Handler{
		BLayer: bLayer,
	}
}

func handleIdParam(ctx *gin.Context) (int, error){
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil{
		return 0, err
	}

	return id, nil
}

func handleError(ctx *gin.Context, err error){
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
}


func (h *Handler) AddUserHandler(ctx *gin.Context){
	var request models.Request	

	err := ctx.ShouldBindJSON(&request)

	handleError(ctx, err)

	err = h.BLayer.AddUser(request)

	handleError(ctx, err)

	ctx.Status(http.StatusCreated)
}

func (h *Handler) GetUsersHandler(ctx *gin.Context){
	users := h.BLayer.GetAllUsers()

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *Handler) GetUserHandler(ctx *gin.Context){
	id, err := handleIdParam(ctx)

	handleError(ctx, err)

	user, err := h.BLayer.GetUser(id)

	handleError(ctx, err)

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUserHandler(ctx *gin.Context){
	id, err := handleIdParam(ctx)

	handleError(ctx, err)

	err = h.BLayer.DeleteUser(id)

	handleError(ctx, err)

	ctx.Status(http.StatusOK)
}

func (h *Handler) UpdateUserHandler(ctx *gin.Context){
	var userRequest models.Request
	id, err := handleIdParam(ctx)
	handleError(ctx, err)

	err = ctx.ShouldBindJSON(&userRequest)
	handleError(ctx, err)

	err = h.BLayer.UpdateUser(userRequest, id)
	handleError(ctx, err)

	ctx.Status(http.StatusOK)
}
