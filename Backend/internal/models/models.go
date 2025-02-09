package models

import "time"

type Container struct {
	ContainerIP     string    `json:"containerIP"`
	PingTimeMKs     int       `json:"pingTimeMKs"`
	LastSuccessDate time.Time `json:"lastSuccessDate"`
}
