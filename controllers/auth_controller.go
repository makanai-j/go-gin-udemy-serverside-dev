package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignUp(ctx *gin.Context)
	LogIn(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var input dto.SignUpInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid SignUp propaty"})
		return
	}

	err := c.service.SignUp(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gailed to create user"})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c *AuthController) LogIn(ctx *gin.Context) {
	var input dto.LogInInput
	if err := ctx.ShouldBindJSON((&input)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.LogIn(input.Email, input.Password)
	if err != nil {
		if err.Error() == "User not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
