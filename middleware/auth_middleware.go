package middleware

import (
    "github.com/gin-gonic/gin"
    "task-manager/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, err := utils.ValidateJWT(c.GetHeader("Authorization"))
        if err != nil {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Set("user_id", userID)
        c.Next()
    }
}
