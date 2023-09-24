package mysql

import (
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Initialize(env env.EnvConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(env.DBDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Printf("banco de dados %s conectado com sucesso", db.Name())
	return db
}
