package server

import (
	"net/http"

	"dsync/logger"
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Run()
}

type server struct {
	ginServer *gin.Engine
	logger    logger.ILogger
}

func NewHttpServer(logger logger.ILogger) IServer {
	newServer := server{
		logger: logger,
	}
	r := gin.Default()
	newServer.logger.Info("running http://localhost:8080")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	newServer.ginServer = r

	return &newServer
}

func (server server) Run() {
	err := server.ginServer.Run()
	if err != nil {
		server.logger.Error(err)
		return
	}
}
