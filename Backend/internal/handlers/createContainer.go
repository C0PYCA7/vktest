package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"vk/Backend/internal/storage"
)

type CreateRequest struct {
	ContainerIP string `json:"containerIP"`
}

type CreateResponse struct {
	Msg   string `json:"msg"`
	Error string `json:"error,omitempty"`
}
type Creater interface {
	CreateContainer(containerIp string) error
}

func CreateContainerHandler(log *slog.Logger, creater Creater) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateRequest

		if err := c.BindJSON(&req); err != nil {
			log.Error("bind json: ", err)
			c.JSON(http.StatusInternalServerError, CreateResponse{
				Msg:   "Failed to create container",
				Error: "bind json",
			})
			return
		}

		err := creater.CreateContainer(req.ContainerIP)
		if err != nil {
			if errors.Is(err, storage.ErrUniqueIP) {
				log.Info("Container with this ip already exists")
				c.JSON(http.StatusBadRequest, CreateResponse{
					Msg:   "Container with this ip already exists",
					Error: "Container already exists",
				})
				return
			}
			if errors.Is(err, storage.ErrBeginTx) {
				log.Error("Begin tx: ", err)
				c.JSON(http.StatusInternalServerError, CreateResponse{
					Msg:   "Don't create container, try again later",
					Error: "Failed to begin tx",
				})
				return
			}
			log.Error("create container: ", err)
			c.JSON(http.StatusInternalServerError, CreateResponse{
				Msg:   "Don't create container, try again later",
				Error: "Failed to create",
			})
			return
		}

		c.JSON(http.StatusOK, CreateResponse{Msg: "Container created"})
		return
	}
}
