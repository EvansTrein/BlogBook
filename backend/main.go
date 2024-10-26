package main

import (
	"backend/envs"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Привет, мир!",
		})
	})

	r.Run(":" + envs.ServerEnvs.BACKEND_PORT)
}
