package msyql

import (
	domain "github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/order"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/database/mysql"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type MysqlAdapterSuiteTest struct {
	suite.Suite
	db  *gorm.DB
	env *env.EnvConfig
}

func (suite *MysqlAdapterSuiteTest) SetupSuite() {
	suite.env = env.LoadConfig("./../../../../test.env")
	suite.db = mysql.Initialize(*suite.env)
	mysql.MigrationUp(*suite.env)
}

func (suite *MysqlAdapterSuiteTest) TearDownTest() {
	suite.db.Table("orders").Where("id is not null").Delete(nil)
}

func (suite *MysqlAdapterSuiteTest) TearDownSuite() {
	mysql.MigrationDown(*suite.env)
}

func (suite *MysqlAdapterSuiteTest) Test_deve_salvar_um_order_no_banco_de_dados() {
	repository := mysql.NewOrderRepository(suite.db)

	order, err := domain.NewOrder("b25635b6-e085-49c3-87fc-71c32fdbb72f", 10.00, 2.00)
	suite.Nil(err)

	err = repository.Save(order)
	suite.Nil(err)

	orderFound, err := repository.FindById(order.ID)
	suite.Nil(err)
	suite.Equal(orderFound, order)

	err = repository.Delete(order.ID)
	suite.Nil(err)
}

func (suite *MysqlAdapterSuiteTest) Test_deve_listar_todos_os_orders() {
	repository := mysql.NewOrderRepository(suite.db)

	order1, err := domain.NewOrder("b25635b6-e085-49c3-87fc-71c32fdbb72f", 10.00, 2.00)
	suite.Nil(err)
	err = repository.Save(order1)
	suite.Nil(err)
	order2, err := domain.NewOrder("7845d2f3-abd5-4c68-a9db-eb71d6622a9a", 10.00, 2.00)
	suite.Nil(err)
	err = repository.Save(order2)
	suite.Nil(err)
	order3, err := domain.NewOrder("e71fdff9-4193-4f92-be57-6e9823966668", 10.00, 2.00)
	suite.Nil(err)
	err = repository.Save(order3)
	suite.Nil(err)

	orders, err := repository.FindAll()
	suite.Nil(err)
	suite.Len(orders, 3)
}

func Test_MysqlAdapterSuiteTest(t *testing.T) {
	suite.Run(t, new(MysqlAdapterSuiteTest))
}
