package domain

type OrderRepository interface {
	Save(entity *Order) error
	FindAll() ([]*Order, error)
}
