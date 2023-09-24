package mysql

import (
	"database/sql"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/infrastructure/env"
	"github.com/pressly/goose/v3"
	"log"
)

func MigrationUp(env env.EnvConfig) {
	gooseDB := getSql(env)

	err := goose.Up(gooseDB, env.MigrationDir)
	if err != nil {
		log.Fatal("Erro ao executar as migrações:", err)
		panic(err)
	}
}

func MigrationDown(env env.EnvConfig) {
	gooseDB := getSql(env)
	err := goose.DownTo(gooseDB, env.MigrationDir, 0)
	if err != nil {
		log.Fatal("Erro ao executar as migrações:", err)
		panic(err)
	}
}

func getSql(env env.EnvConfig) *sql.DB {
	gooseDB, err := goose.OpenDBWithDriver(env.DBDriver, env.DBDsn)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
		panic(err)
	}

	err = goose.SetDialect(env.DBDriver)
	if err != nil {
		log.Fatal("Erro ao configurar o dialect:", err)
		panic(err)
	}
	return gooseDB
}
