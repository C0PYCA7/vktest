package workerpool

import (
	"log"
	"sync"
	"vk/Pinger/internal/models"
	"vk/Pinger/internal/pinger"
)

type Sender interface {
	SendMessage(topic string, data models.Container) error
}

func WorkerPool(containers []models.Container, sender Sender, topic string) {
	var (
		workerPool = 2
		result     = make(chan *models.Container, len(containers))
		jobs       = make(chan models.Container, len(containers))
		wgWorker   = sync.WaitGroup{}
		wgListener = sync.WaitGroup{}
	)

	for w := 1; w <= workerPool; w++ {
		wgWorker.Add(1)
		go func(w int) {
			defer wgWorker.Done()
			worker(jobs, result, w)
		}(w)
	}

	wgListener.Add(1)
	go func() {
		defer wgListener.Done()

		for value := range result {
			err := sender.SendMessage(topic, *value)
			if err != nil {
				log.Println("send message to kafka: ", err)
			}
		}
	}()

	for i := 0; i < len(containers); i++ {
		jobs <- containers[i]
	}

	close(jobs)
	wgWorker.Wait()

	close(result)
	wgListener.Wait()
}

func worker(jobs chan models.Container, result chan *models.Container, id int) {
	for j := range jobs {
		container, _ := pinger.PingContainer(j)
		log.Println("контейнер в воркере после пинга: ", container)
		result <- container
	}
	log.Printf("работник %d закончил работать\n", id)
}
