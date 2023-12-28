package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Start Request")
		ctx.Next()
		fmt.Println("End Request")
	}
}
