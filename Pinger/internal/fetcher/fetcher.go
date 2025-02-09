package fetcher

import (
	"encoding/json"
	"log"
	"net/http"
	"vk/Pinger/internal/models"
)

func GetContainers(addr string) []models.Container {
	resp, err := http.Get(addr)
	if err != nil {
		log.Println("failed to get containers: ", err)
		return nil
	}

	var containers models.Containers

	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		log.Println("failed to decode containers:", err)
		return nil
	}
	return containers.Containers
}
