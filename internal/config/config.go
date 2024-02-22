package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ApnsConfig struct {
	KeyID       string `env:"KEY_ID"`
	TeamID      string `env:"TEAM_ID"`
	KeyFile     string `env:"KEY_FILE"`
	DeviceToken string `env:"DEVICE_TOKEN"`
	Topic       string `env:"TOPIC"`
}

type Config struct {
	Apns ApnsConfig
}

func Get() *Config {
	var cfg Config
	//nolint:errcheck
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Println("Error reading env:", err.Error())
		// panic if can't read env will prevent the app from running
		// better than spawning a new pod with missing env
		panic(err)
	}
	return &cfg
}
