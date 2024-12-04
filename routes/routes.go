package routes

import (
    "github.com/gin-gonic/gin"
    "task-manager/controllers"
    "task-manager/middleware"
)

func RegisterRoutes(router *gin.Engine) {
    auth := router.Group("/auth")
    {
        auth.POST("/register", controllers.RegisterUser)
        auth.POST("/login", controllers.LoginUser)
    }

    tasks := router.Group("/tasks").Use(middleware.AuthMiddleware())
    {
        tasks.POST("/", controllers.CreateTask)
    }

    user := router.Group("/users")
    {
        RegisterUserRoutes(user)
    }
}
