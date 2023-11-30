package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id         uuid.UUID   `json:"id"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"orderItems"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}
