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

func ProcessIncomingOrder(order model.Order) {
	log.Println("Process incoming order: ", order.Id, " in ", order.PrepTime)
	payload := common.UpdateOrderRequestDto{
		Id:     order.Id,
		Status: common.OrderStatus["InProgress"],
	}

	log.Println("Order In Progress...")

	UpdateOrder(payload)

	time.Sleep(order.PrepTime)

	payload = common.UpdateOrderRequestDto{
		Id:     order.Id,
		Status: common.OrderStatus["Ready"],
	}

	UpdateOrder(payload)

	log.Printf("OrderId %v Is Done!\n", order.Id)
}
