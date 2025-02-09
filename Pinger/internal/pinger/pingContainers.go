package pinger

import (
	"github.com/go-ping/ping"
	"log"
	"time"
	"vk/Pinger/internal/models"
)

func PingContainer(container models.Container) (*models.Container, error) {
	pinger, err := ping.NewPinger(container.ContainerIP)
	if err != nil {
		log.Println("Create pinger: ", err)
		return nil, err
	}
	pinger.Count = 4
	pinger.Timeout = 5 * time.Second
	pinger.SetPrivileged(true)
	err = pinger.Run()
	if err != nil {
		log.Println("Start pinger: ", err)
		return nil, err
	}
	stats := pinger.Statistics()

	if stats.AvgRtt.Microseconds() == 0 {
		log.Println("Container doesn't work")
		return &models.Container{
			ContainerIP:     container.ContainerIP,
			PingTimeMks:     container.PingTimeMks,
			LastSuccessDate: container.LastSuccessDate,
		}, nil
	}
	return &models.Container{
		ContainerIP:     container.ContainerIP,
		PingTimeMks:     stats.AvgRtt.Microseconds(),
		LastSuccessDate: time.Now(),
	}, nil
}
