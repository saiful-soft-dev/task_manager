package controllers

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    // "gorm.io/gorm"
    "task-manager/database"
    "task-manager/models"
    "task-manager/utils"
)

func RegisterUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(201, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    var user models.User
    if err := database.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
        c.JSON(401, gin.H{"error": "Invalid username or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
        c.JSON(401, gin.H{"error": "Invalid username or password"})
        return
    }

    token, _ := utils.GenerateJWT(user.ID)
    c.JSON(200, gin.H{"token": token})
}
