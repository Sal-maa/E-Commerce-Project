package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type OrderRepository interface {
	CreateOrder1(order entity.Order) (entity.Order, error)
	CreateOrder2(order entity.Order) (entity.Order, error)
	CreateOrder3(order entity.Order) (entity.Order, error)
	UpdateOrder(order entity.Order) (entity.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder1(order entity.Order) (entity.Order, error) {
	_, err1 := r.db.Exec(`INSERT INTO orders(created_at, updated_at, cart_id, status_order,order_date,total)
						VALUES(?,?,?,?,?,?);
						`, order.CreatedAt, order.UpdatedAt, order.Cart.Id, order.StatusOrder, order.OrderDate, order.Total)
	return order, err1
}

func (r *orderRepository) CreateOrder2(order entity.Order) (entity.Order, error) {
	_, err2 := r.db.Exec(`INSERT INTO address(street,city,state,zip)
						VALUES(?,?,?,?);
						`, order.Address.Street, order.Address.City, order.Address.State, order.Address.Zip)
	return order, err2
}

func (r *orderRepository) CreateOrder3(order entity.Order) (entity.Order, error) {
	_, err3 := r.db.Exec(`INSERT INTO creditcarts(type,name,number,cvv)
						VALUES(?,?,?,?)
						`, order.CreditCard.Type, order.CreditCard.Name, order.CreditCard.Number, order.CreditCard.CVV)
	return order, err3
}


func (r *orderRepository) UpdateOrder(order entity.Order) (entity.Order, error) {
	_, err := r.db.Exec("UPDATE orders SET updated_at = ?, status_order = ? WHERE id = ?", order.UpdatedAt, order.StatusOrder, order.Id)
	if err != nil {
		fmt.Println("update error:", err)
	}
	return order, err
}