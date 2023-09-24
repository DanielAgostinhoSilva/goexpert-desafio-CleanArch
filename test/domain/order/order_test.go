package order

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OrderSuitTest struct {
	suite.Suite
}

func (suite *OrderSuitTest) Test_deve_lancar_um_erro_quando_o_id_estiver_invalido() {
	order, err := domain.NewOrder("", 10.00, 10.00)
	suite.Nil(order)
	suite.ErrorContains(err, "invalid id")
}

func (suite *OrderSuitTest) Test_deve_lancar_um_erro_quando_o_price_estiver_invalido() {
	order, err := domain.NewOrder("123", 0, 10.00)
	suite.Nil(order)
	suite.ErrorContains(err, "invalid price")
}

func (suite *OrderSuitTest) Test_deve_lancar_um_erro_quando_o_tax_estiver_invalido() {
	order, err := domain.NewOrder("123", 10.00, -32)
	suite.Nil(order)
	suite.ErrorContains(err, "invalid tax")
}

func (suite *OrderSuitTest) Test_deve_criar_um_order_valido() {
	order, err := domain.NewOrder("123", 10.00, 2.0)
	suite.Nil(err)
	suite.Equal("123", order.ID)
	suite.Equal(10.00, order.Price)
	suite.Equal(2.0, order.Tax)
}

func (suite *OrderSuitTest) Test_deve_calcular_o_valor_do_price_mais_o_tax() {
	order, err := domain.NewOrder("123", 10.00, 2.0)
	suite.Nil(err)
	err = order.CalculateFinalPrice()
	suite.Nil(err)
	suite.Equal(12.00, order.FinalPrice)
}

func Test_OrderSuitTest(t *testing.T) {
	suite.Run(t, new(OrderSuitTest))
}
