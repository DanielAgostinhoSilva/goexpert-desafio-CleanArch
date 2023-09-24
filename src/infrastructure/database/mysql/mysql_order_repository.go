package mysql

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"gorm.io/gorm"
)

type MysqlOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *MysqlOrderRepository {
	return &MysqlOrderRepository{db: db}
}

func (props *MysqlOrderRepository) Save(order *domain.Order) error {
	return props.db.Create(order).Error
}

func (props *MysqlOrderRepository) FindAll() ([]domain.Order, error) {
	var models []domain.Order
	err := props.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, err
}

func (props *MysqlOrderRepository) FindById(id string) (*domain.Order, error) {
	return props.findOrFail(id)
}

func (props *MysqlOrderRepository) Delete(id string) error {
	order, err := props.findOrFail(id)
	if err != nil {
		return err
	}
	return props.db.Delete(order).Error
}

func (props *MysqlOrderRepository) findOrFail(id string) (*domain.Order, error) {
	var order domain.Order
	err := props.db.First(&order, "id = ?", id).Error
	return &order, err
}

func (props *MysqlOrderRepository) GetTotal() (int64, error) {
	var total int64
	err := props.db.Table("orders").Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}
