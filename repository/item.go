package repository

import (
	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/util"
)

func FindAllItems() []model.Item {
	var items []model.Item
	util.OpenFile("items.json", &items)
	return items
}
