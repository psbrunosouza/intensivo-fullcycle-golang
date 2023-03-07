package database

import (
	"database/sql"
	"tax-app/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, error := r.Db.Exec("insert into orders (id, price, tax, final_price) values (?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if error != nil {
		return error
	}

	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("select count(*) from orders").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
