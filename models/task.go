package models

type Task struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Title       string `json:"title"`
    Description string `json:"description"`
    IsCompleted bool   `json:"is_completed"`
    UserID      uint   `json:"user_id"` // Foreign key
}
