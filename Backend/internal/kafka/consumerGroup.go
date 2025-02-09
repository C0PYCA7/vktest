package kafka

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
	"vk/Backend/internal/models"
	"vk/Backend/internal/storage"
)

type Updater interface {
	UpdateContainerInfo(containerIP string, pingTimeMs int, lastSuccessData time.Time) error
}

type Consumer struct {
	Ready   chan bool
	Updater Updater
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.Ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, partition = %d", string(message.Value), message.Timestamp, message.Topic, message.Partition)
			var container models.Container
			err := json.Unmarshal(message.Value, &container)
			if err != nil {
				log.Println("unmarshal message in consume: ", err)
				return err
			}
			err = consumer.Updater.UpdateContainerInfo(container.ContainerIP, container.PingTimeMKs, container.LastSuccessDate)
			if err != nil {
				if errors.Is(err, storage.ErrBeginTx) {
					return fmt.Errorf("%v", storage.ErrBeginTx)
				}
				return err
			}
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			log.Println("Session context done, rebalance might be happening")
			return nil
		}
	}
}
