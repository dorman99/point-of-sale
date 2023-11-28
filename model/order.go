package model

import "github.com/google/uuid"

type Order struct {
	Id         uuid.UUID   `json:"id"`
	Total      int         `json:"total"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"orderItems"`
}
