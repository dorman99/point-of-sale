package service

import (
	"time"

	"github.com/dorman99/point-of-sales/common"
	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/repository"
	"github.com/dorman99/point-of-sales/util"
)

func CreateOrder(orderItems []model.OrderItem) model.Order {
	orders := repository.FindAllOrders()

	// caculate order items
	order := model.Order{
		Id:         util.GenerateUUID(),
		Status:     string(common.OrderStatus["Todo"]),
		OrderItems: orderItems,
		PrepTime:   calculatePrepTime(orderItems),
		Total:      calculateOrderItems(orderItems),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	orders = append(orders, order)

	repository.UpdateOrders(orders)

	return order
}

func calculateOrderItems(orderItems []model.OrderItem) int {
	res := 0
	for _, o := range orderItems {
		res += (o.Quantity * o.Item.Price)
	}
	return res
}

func calculatePrepTime(oi []model.OrderItem) time.Duration {
	secondsPerItems := 5
	totalInSeconds := secondsPerItems * len(oi)
	return time.Duration(time.Second * time.Duration(totalInSeconds))
}
