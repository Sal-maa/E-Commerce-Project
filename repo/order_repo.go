package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Sal-maa/E-Commerce-Project/entity"
)

type OrderRepository interface {
	CreateAddress(order entity.Order) (entity.Address, error)
	CreatePayment(order entity.Order) (entity.CreditCard, error)
	CreateOrder(order entity.Order) (int, error)
	CreateOrderDetail(orderDetail entity.CreateOrderDetailRequest) (entity.CreateOrderDetailRequest, error)
	GetId() (int, error)
	GetOrder(id int) ([]entity.Order, error)
	UpdateOrder(order entity.Order) (entity.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateAddress(order entity.Order) (entity.Address, error) {
	_, err := r.db.Exec(`INSERT INTO address(street,city,state,zip)
						VALUES(?,?,?,?);
						`, order.Address.Street, order.Address.City, order.Address.State, order.Address.Zip)

	result, err1 := r.db.Query(`SELECT id FROM address WHERE street=? AND city=? AND state=? AND zip=? ORDER BY id DESC LIMIT 1
							`, order.Address.Street, order.Address.City, order.Address.State, order.Address.Zip)
	if err1 != nil {
		fmt.Println("failed in query", err1)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	address := entity.Address{}
	errScan := result.Scan(&address.Id)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}
	return address, err
}

func (r *orderRepository) CreatePayment(order entity.Order) (entity.CreditCard, error) {
	_, err := r.db.Exec(`INSERT INTO creditcarts(type,name,number,cvv)
						VALUES(?,?,?,?)
						`, order.CreditCard.Type, order.CreditCard.Name, order.CreditCard.Number, order.CreditCard.CVV)

	result, err1 := r.db.Query(`SELECT id FROM creditcarts WHERE type=? AND name=? AND number=? AND cvv=? ORDER BY id DESC LIMIT 1
						`, order.CreditCard.Type, order.CreditCard.Name, order.CreditCard.Number, order.CreditCard.CVV)
	if err1 != nil {
		fmt.Println("failed in query", err1)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	payment := entity.CreditCard{}
	errScan := result.Scan(&payment.Id)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}
	return payment, err
}

func (r *orderRepository) CreateOrder(order entity.Order) (int, error) {
	cartByte, _ := json.Marshal([]int(order.Cart))
	cartString := string(cartByte)
	_, err := r.db.Exec(`INSERT INTO orders(created_at, updated_at, user_id, address_id, creditcard_id, cart_id, status_order, order_date, total)
						VALUES(?,?,?,?,?,?,?,?,?);
						`, order.CreatedAt, order.UpdatedAt, order.User.Id, order.Address.Id, order.CreditCard.Id, cartString, order.StatusOrder, order.OrderDate, order.Total)

	result, errId := r.db.Query(`SELECT id FROM orders WHERE created_at = ? AND updated_at = ?  AND user_id = ? 
								AND address_id = ? AND creditcard_id = ? AND cart_id = ? AND status_order = ? AND order_date = ? AND total = ? ORDER BY id DESC LIMIT 1`,
		order.CreatedAt, order.UpdatedAt, order.User.Id, order.Address.Id, order.CreditCard.Id, cartString, order.StatusOrder, order.OrderDate, order.Total)
	if errId != nil {
		fmt.Println("failed in query", errId)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	var orderId int
	errScan := result.Scan(&orderId)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}
	return orderId, err
}

func (r *orderRepository) CreateOrderDetail(orderDetail entity.CreateOrderDetailRequest) (entity.CreateOrderDetailRequest, error) {
	_, err := r.db.Exec(`INSERT INTO order_details(order_id, cart_id) VALUES(?,?)`, orderDetail.OrderId, orderDetail.CartId)
	if err != nil {
		return orderDetail, err
	}
	return orderDetail, nil
}

func (r *orderRepository) GetId() (int, error) {
	result, err := r.db.Query(`SELECT id FROM orders ORDER BY id DESC LIMIT 1`)
	if err != nil {
		fmt.Println("failed in query", err)
	}

	defer result.Close()

	if isExist := result.Next(); !isExist {
		fmt.Println("data not found", err)
	}
	var id int
	errScan := result.Scan(&id)
	if errScan != nil {
		fmt.Println("failed to read data", errScan)
	}
	return id, err
}

func (r *orderRepository) UpdateOrder(order entity.Order) (entity.Order, error) {
	_, err := r.db.Exec("UPDATE orders SET updated_at = ?, status_order = ? WHERE id = ?", order.UpdatedAt, order.StatusOrder, order.Id)
	if err != nil {
		fmt.Println("update error:", err)
	}
	return order, err
}

func (r *orderRepository) GetOrder(id int) ([]entity.Order, error) {
	orders := []entity.Order{}
	result, err := r.db.Query(`SELECT`)
}
