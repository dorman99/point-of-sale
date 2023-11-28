package service

import (
	"github.com/dorman99/point-of-sales/model"
	"github.com/dorman99/point-of-sales/repository"
)

func FindAllItems() []model.Item {
	return repository.FindAllItems()
}
