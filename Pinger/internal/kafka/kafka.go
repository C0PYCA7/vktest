package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"vk/Pinger/internal/models"
)

type AsyncProducer struct {
	producer sarama.AsyncProducer
}

func NewAsyncProducer(addr string) (*AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	asyncProducer, err := sarama.NewAsyncProducer([]string{addr}, config)
	if err != nil {
		log.Println("Create async producer: ", err)
		return nil, err
	}

	go func() {
		for {
			select {
			case success := <-asyncProducer.Successes():
				log.Printf("Message sent successfully: value = %s, topic = %s, partition = %d, offset = %d",
					success.Value, success.Topic, success.Partition, success.Offset)
			case err := <-asyncProducer.Errors():
				log.Printf("Send message: %v", err)
			}
		}
	}()

	return &AsyncProducer{producer: asyncProducer}, nil
}

func CheckTopicExists(admin sarama.ClusterAdmin, topicName string) (bool, error) {
	topics, err := admin.ListTopics()
	if err != nil {
		return false, err
	}

	_, exists := topics[topicName]
	if exists {
		return true, nil
	}
	return false, nil
}

func NewTopic(addr, topicName string, numPartitions int32, replica int16) error {
	admin, err := sarama.NewClusterAdmin([]string{addr}, sarama.NewConfig())
	if err != nil {
		return err
	}
	defer admin.Close()

	exists, err := CheckTopicExists(admin, topicName)
	if err != nil {
		return err
	}

	if exists {
		log.Println("topic already exists")
		return nil
	}

	retentionMs := "604800000"

	topicDetail := sarama.TopicDetail{
		NumPartitions:     numPartitions,
		ReplicationFactor: replica,
		ConfigEntries: map[string]*string{
			"retention.ms": &retentionMs,
		},
	}
	err = admin.CreateTopic(topicName, &topicDetail, false)
	if err != nil {
		return err
	}
	return nil
}

func (a *AsyncProducer) SendMessage(topic string, data models.Container) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(dataBytes),
	}

	a.producer.Input() <- message
	log.Println("Сообщение отправлено: ", message)
	return nil
}
