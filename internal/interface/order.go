package model

import (
	"tax-app/internal/entity"
)

type Order interface {
	Save(order *entity.Order) error
	GetTotal() (int, error)
}
