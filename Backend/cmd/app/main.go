package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log2 "log"
	"os"
	"vk/Backend/internal/config"
	"vk/Backend/internal/handlers"
	"vk/Backend/internal/kafka"
	"vk/Backend/internal/logger"
	"vk/Backend/internal/storage/postgres"
)

func main() {
	cfg := config.New()
	log2.Println(cfg)
	db := postgres.New(cfg.Database.Dsn)
	log := logger.New()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.Frontend.Addr},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	group, err := kafka.NewConsumerGroup(cfg.Kafka.Port, cfg.Kafka.GroupId)
	if err != nil {
		log.Error("error: ", err)
		os.Exit(1)
	}
	go group.StartListening(cfg.Kafka.Topic, db)
	router.GET("/", handlers.GetAllContainersDataHandler(log, db))
	router.POST("/create", handlers.CreateContainerHandler(log, db))

	router.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
