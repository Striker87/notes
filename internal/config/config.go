package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"true"`
	Listen        struct {
		Type       string `env:"LISTEN_TYPE" env-default:"port" env-description:"port or socket. if socket, then specify SOCKET_FILE is required"`
		BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port       string `env:"Port" env-default:"8080"`
		SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
	}
	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `env:"ADMIN_PWD" env-default:"admin"`
		}
	}
	PostgreSQL struct {
		Username string `env:"PSQL_USERNAME" env-required:"true"`
		Password string `env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `env:"PSQL_HOST" env-required:"true"`
		Port     string `env:"PSQL_PORT" env-required:"true"`
		Database string `env:"PSQL_DATABASE" env-required:"true"`
	}
}

func NewConfig() *Config {
	config := &Config{}
	if err := cleanenv.ReadEnv(config); err != nil {
		helpText := "Monolith notes system"
		help, _ := cleanenv.GetDescription(config, &helpText)
		log.Println(help)
		log.Fatal(err)
	}
	return config
}
