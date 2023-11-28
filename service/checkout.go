package service

import (
	"log"

	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/repository"
	"github.com/dorman99/point-of-sales/util"
)

func CreateOrder(orderItems []model.OrderItem) []model.Order {
	orders := repository.FindAllOrders()

	// caculate order items
	order := model.Order{
		Id:         util.GenerateUUID(),
		Status:     "ToDo",
		OrderItems: orderItems,
		Total:      calculateOrderItems(orderItems),
	}

	log.Println(orders)

	orders = append(orders, order)

	repository.UpdateOrders(orders)

	return orders
}

func calculateOrderItems(orderItems []model.OrderItem) int {
	res := 0
	for _, o := range orderItems {
		res += (o.Quantity * o.Item.Price)
	}
	return res
}
