package env

import (
	"github.com/spf13/viper"
	"log"
)

type EnvConfig struct {
	MigrationDir      string `mapstructure:"MIGRATION_DIR"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBDsn             string `mapstructure:"DB_DSN"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(filePath string) *EnvConfig {
	var cfg *EnvConfig
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	log.Println("arquivo .env carregado")
	return cfg
}
