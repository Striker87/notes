package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"true"`
	Listen        struct {
		Type   string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"Port" env-default:"8080"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-required:"true"`
			Password string `env:"ADMIN_PWD" env-required:"true"`
		}
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
