package common

import "github.com/google/uuid"

type UpdateOrderRequestDto struct {
	Id     uuid.UUID    `json:"id"`
	Status TOrderStatus `json:"status"`
}

type ResponseData struct {
	Success string      `json:"success"`
	Data    interface{} `json:"data"`
}
