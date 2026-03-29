package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Cfg struct {
	Service Service  `yaml:"service"`
	Queue   RabbitMQ `yaml:"queue"`
}

type Service struct {
	Host     string `yaml:"host"`
	MainPort string `yaml:"mainPort"`
}

type RabbitMQ struct {
	RabbitMQURL string `yaml:"rabbitMQURL"`
	QueueName   string `yaml:"queueName"`
}

func (cfg *Cfg) MustLoad() {
	confPath := os.Getenv("CONFIG_FILE")
	log.Printf("Config path: %s", confPath)
	viper.SetConfigFile(confPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error to read file: %w", err)
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("error to decode file: %w", err)
	}
}
