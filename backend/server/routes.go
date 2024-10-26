package server

import (
	"backend/envs"
	"backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRotes() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", handlers.GetHandler)

	router.Run(":" + envs.ServerEnvs.BACKEND_PORT)
}
