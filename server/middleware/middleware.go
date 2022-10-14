package middleware

import (
	"final-project/server/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization")
	tokenArr := strings.Split(tokenHeader, "Bearer ")

	if len(tokenArr) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	tokenStr := tokenArr[1]

	payload, err := helper.ValidateToken(tokenStr)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Set("email", payload["email"])
	ctx.Next()
}
