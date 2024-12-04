package routes

import (
    "github.com/gin-gonic/gin"
    "task-manager/controllers"
    "task-manager/middleware"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
    router.GET("/me", middleware.AuthMiddleware(), controllers.GetUser)     // Get user details
    router.PUT("/me", middleware.AuthMiddleware(), controllers.UpdateUser) // Update user details
}
