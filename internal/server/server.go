package server

import (
	"github.com/alejolencinas/dollar-price/internal/api"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", api.Ping)
		v1.GET("/dollar", api.GetDollarPrice)
	}

	return r
}
