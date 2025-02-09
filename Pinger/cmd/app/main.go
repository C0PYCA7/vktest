package main

import (
	"log"
	"time"
	"vk/Pinger/internal/config"
	"vk/Pinger/internal/fetcher"
	"vk/Pinger/internal/kafka"
	"vk/Pinger/internal/workerpool"
)

func main() {
	cfg := config.New()
	err := kafka.NewTopic(cfg.Kafka.Port, cfg.Kafka.Topic, 3, 1)
	if err != nil {
		log.Println(err)
	}

	asyncProducer, err := kafka.NewAsyncProducer(cfg.Kafka.Port)
	if err != nil {
		log.Println(err)
	}

	containers := fetcher.GetContainers(cfg.Server.Addr)
	log.Println(containers)
	workerpool.WorkerPool(containers, asyncProducer, cfg.Kafka.Topic)
	time.Sleep(time.Millisecond * 5)
}
