package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ Could not load .env file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("❌ Failed to unmarshal config: %v", err)
	}

	fmt.Println("✅ Configuration loaded")
	return &cfg
}
