package domain

type OrderRepository interface {
	Save(entity *Order) error
	FindAll() ([]Order, error)
	FindById(id string) (*Order, error)
	Delete(id string) error
}
