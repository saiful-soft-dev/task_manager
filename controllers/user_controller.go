package controllers

import (
    "github.com/gin-gonic/gin"
    "task-manager/utils"  // Make sure this import path is correct
    "task-manager/models"
    "task-manager/database"
    // "net/http"
)

// GetUser retrieves the details of the currently logged-in user.
func GetUser(c *gin.Context) {
    userID := c.GetUint("user_id")

    var user models.User
    if err := database.DB.Preload("Tasks").First(&user, userID).Error; err != nil {
        c.JSON(404, gin.H{"error": "User not found"})
        return
    }

    c.JSON(200, user)
}

// UpdateUser allows the logged-in user to update their details.
func UpdateUser(c *gin.Context) {
    userID := c.GetUint("user_id")

    var user models.User
    if err := database.DB.First(&user, userID).Error; err != nil {
        c.JSON(404, gin.H{"error": "User not found"})
        return
    }

    var updatedData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    if updatedData.Username != "" {
        user.Username = updatedData.Username
    }

    if updatedData.Password != "" {
        hashedPassword, err := utils.HashPassword(updatedData.Password)
        if err != nil {
            c.JSON(500, gin.H{"error": "Failed to hash password"})
            return
        }
        user.Password = hashedPassword
    }

    if err := database.DB.Save(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(200, gin.H{"message": "User updated successfully"})
}
