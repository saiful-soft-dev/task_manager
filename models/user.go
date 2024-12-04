package models

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"-"` // Hashed password
    Tasks    []Task `json:"tasks"` // One-to-many relationship
}
