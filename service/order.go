package service

import (
	"log"
	"time"

	"github.com/dorman99/point-of-sales/common"
	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/repository"
	"github.com/google/uuid"
)

func UpdateOrder(payload common.UpdateOrderRequestDto) {
	orders := repository.FindAllOrders()
	var order model.Order
	var ordIdx int

	for i, o := range orders {
		if o.Id == payload.Id {
			order = o
			ordIdx = i
		}
	}

	if order.Id == uuid.Nil {
		log.Println("Order Id Not found ", payload.Id)
		return
	}

	order.Status = string(payload.Status)
	order.UpdatedAt = time.Now()
	orders[ordIdx] = order

	repository.UpdateOrders(orders)
}
