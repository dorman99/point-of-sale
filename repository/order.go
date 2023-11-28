package repository

import (
	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/util"
)

func FindAllOrders() []model.Order {
	var orders []model.Order
	util.OpenFile("orders.json", &orders)
	return orders
}

func UpdateOrders(orders []model.Order) {
	util.SaveFile("orders.json", orders)
}
