package config

import (
	"github.com/gin-gonic/gin"
)

/*
Server hold the gin.Engine instance
*/
type Server struct {
	Engine *gin.Engine
	Port   string
}

func NewServer(port string) *Server {

	engine := gin.Default()

	return &Server{
		Engine: engine,
		Port:   port,
	}
}
