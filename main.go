package main

import (
    "github.com/gin-gonic/gin"
    "task-manager/routes"
    "task-manager/database"
)

func main() {
    // Initialize the database
    database.Connect()
    database.Migrate()

    // Create the router
    router := gin.Default()

    // Register routes
    routes.RegisterRoutes(router)

    // Start the server
    router.Run(":8080")
}
