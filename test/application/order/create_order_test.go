package application

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CreateOrderUseCaseSuiteTest struct {
	suite.Suite
	OrderRepository *MockOrderRepository
}

func (suite *CreateOrderUseCaseSuiteTest) SetupTest() {
	suite.OrderRepository = new(MockOrderRepository)
}

func (suite *CreateOrderUseCaseSuiteTest) Test_deve_executar_um_CreateOrderUseCase() {
	suite.OrderRepository.On("Save", mock.AnythingOfType("*domain.Order")).Return(nil)
	createOrderUseCase := application.NewCreateOrderUseCase(suite.OrderRepository)

	createOrderOutput, err := createOrderUseCase.Execute(application.CreateOrderInput{
		ID:    "123",
		Price: 10.00,
		Tax:   2.0,
	})

	suite.Nil(err)
	suite.NotNil(createOrderOutput)
	suite.Equal("123", createOrderOutput.ID)
	suite.Equal(10.00, createOrderOutput.Price)
	suite.Equal(2.0, createOrderOutput.Tax)
	suite.Equal(12.0, createOrderOutput.FinalPrice)
}

func Test_CreateOrderUseCaseSuiteTest(t *testing.T) {
	suite.Run(t, new(CreateOrderUseCaseSuiteTest))
}

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(entity *domain.Order) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *MockOrderRepository) FindAll() ([]*domain.Order, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Order), args.Error(1)
}
