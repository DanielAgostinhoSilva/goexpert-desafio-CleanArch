package mocks

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(entity *domain.Order) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockOrderRepository) FindAll() ([]domain.Order, error) {
	args := m.Called()
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderRepository) FindById(id string) (*domain.Order, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
