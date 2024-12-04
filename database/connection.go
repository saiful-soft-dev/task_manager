package database

import (
    "log"
    "task-manager/config"
    "task-manager/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    config.LoadConfig()

    dsn := "host=" + config.GetEnv("DB_HOST") +
        " user=" + config.GetEnv("DB_USER") +
        " password=" + config.GetEnv("DB_PASSWORD") +
        " dbname=" + config.GetEnv("DB_NAME") +
        " port=5432 sslmode=disable"

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
}

func Migrate() {
    DB.AutoMigrate(&models.User{}, &models.Task{})
}
