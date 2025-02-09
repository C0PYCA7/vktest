package models

import "time"

type Containers struct {
	Containers []Container `json:"containers"`
}
type Container struct {
	ContainerIP     string    `json:"containerIP"`
	PingTimeMks     int64     `json:"pingTimeMKs"`
	LastSuccessDate time.Time `json:"lastSuccessDate"`
}
