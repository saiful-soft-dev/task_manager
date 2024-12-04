package controllers

import (
    "github.com/gin-gonic/gin"
    "task-manager/database"
    "task-manager/models"
)

func CreateTask(c *gin.Context) {
    userID := c.GetUint("user_id")

    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    task.UserID = userID
    if err := database.DB.Create(&task).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create task"})
        return
    }

    c.JSON(201, task)
}
