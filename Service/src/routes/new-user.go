package routes

import "github.com/gin-gonic/gin"

func NewUserHandler(ctx *gin.Context) {
	// Handle new user creation
	ctx.String(200, "New User Created")
}
