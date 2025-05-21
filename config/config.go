package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"telegram_token"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig() // игнорируем ошибку, если файла нет, нужно будет проверить

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
