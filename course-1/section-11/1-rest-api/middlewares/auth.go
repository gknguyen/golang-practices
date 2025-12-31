package middlewares

import (
	"net/http"
	"restApi/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
