package handlers

import (
	"github.com/gin-gonic/gin"
	log2 "log"
	"log/slog"
	"net/http"
	"vk/Backend/internal/models"
)

type Selecter interface {
	SelectAllContainersData() ([]models.Container, error)
}

type GetAllContainersDataResponse struct {
	Containers []models.Container `json:"containers"`
	Error      string             `json:"error,omitempty"`
}

func GetAllContainersDataHandler(log *slog.Logger, selecter Selecter) gin.HandlerFunc {
	return func(c *gin.Context) {
		containers, err := selecter.SelectAllContainersData()
		if err != nil {
			log.Error("Get containers data error: ", err)
			c.JSON(http.StatusInternalServerError, GetAllContainersDataResponse{
				Containers: nil,
				Error:      "Can't get containers data",
			})
			return
		}

		log2.Println(containers)
		c.JSON(http.StatusOK, GetAllContainersDataResponse{
			Containers: containers,
		})
		return
	}
}
