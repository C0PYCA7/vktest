package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Kafka    KafkaConfig
	Frontend FrontendConfig
}

type DatabaseConfig struct {
	Dsn string
}

type ServerConfig struct {
	Port string
}

type KafkaConfig struct {
	Port    string
	GroupId string
	Topic   string
}

type FrontendConfig struct {
	Addr string
}

func New() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Load config: %v", err)
		os.Exit(1)
	}
	return &Config{
		Database: DatabaseConfig{
			Dsn: os.Getenv("DB_DSN"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		Kafka: KafkaConfig{
			Port:    os.Getenv("KAFKA_PORT"),
			GroupId: os.Getenv("KAFKA_GROUP_ID"),
			Topic:   os.Getenv("KAFKA_TOPIC"),
		},
		Frontend: FrontendConfig{
			Addr: os.Getenv("FRONT_ADDR"),
		},
	}
}
