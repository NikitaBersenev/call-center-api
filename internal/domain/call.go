package domain

import "time"

type Call struct {
	ID          uint64    `json:"id" gorm:"primary_key"`
	ClientName  string    `json:"client_name" binding:"required"`
	PhoneNumber string    `json:"phone_number"`
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"in=открыта,закрыта"`
	CreatedAt   time.Time `json:"created_at"`
}
