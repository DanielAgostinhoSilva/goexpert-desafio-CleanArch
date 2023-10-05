package application

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/application/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CreateOrderUseCaseSuiteTest struct {
	suite.Suite
	OrderRepository *mocks.MockOrderRepository
	OrderEvent      *mocks.MockEvent
	EventDispatcher *mocks.MockEventDispatcher
}

func (suite *CreateOrderUseCaseSuiteTest) SetupTest() {
	suite.OrderRepository = new(mocks.MockOrderRepository)
	suite.OrderEvent = new(mocks.MockEvent)
	suite.EventDispatcher = new(mocks.MockEventDispatcher)
}

func (suite *CreateOrderUseCaseSuiteTest) Test_deve_executar_um_CreateOrderUseCase() {
	suite.OrderRepository.On("Save", mock.AnythingOfType("*domain.Order")).Return(nil)
	suite.OrderEvent.On("SetPayload", mock.AnythingOfType("application.CreateOrderOutput"))
	suite.EventDispatcher.On("Dispatch", mock.AnythingOfType("*mocks.MockEvent")).Return(nil)

	createOrderUseCase := application.NewCreateOrderUseCase(suite.OrderRepository, suite.OrderEvent, suite.EventDispatcher)

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
