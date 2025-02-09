package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Kafka  KafkaConfig
	Server ServerConfig
}

type ServerConfig struct {
	Addr string
}

type KafkaConfig struct {
	Port    string
	GroupId string
	Topic   string
}

func New() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Load config: %v", err)
		os.Exit(1)
	}
	return &Config{
		Server: ServerConfig{
			Addr: os.Getenv("SERVER_ADDRESS"),
		},
		Kafka: KafkaConfig{
			Port:    os.Getenv("KAFKA_PORT"),
			GroupId: os.Getenv("KAFKA_GROUP_ID"),
			Topic:   os.Getenv("KAFKA_TOPIC"),
		},
	}
}
