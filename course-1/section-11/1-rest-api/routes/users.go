package routes

import (
	"net/http"
	"restApi/models"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.Users

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
