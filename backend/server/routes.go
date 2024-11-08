package server

import (
	"backend/envs"
	"backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRotes() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:7100"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// user create
	router.POST("/user/registration", handlers.SignUpUserHandler)

	// user authorization
	router.POST("/user/login", handlers.SignInUserHandler)

	// update user info
	// router.PUT("/user/:id", handlers.UpdateUeserHandler)

	// getting all users
	router.GET("/users", handlers.AuthMiddleware(), handlers.GetUsersHandler)

	// delete user
	// router.DELETE("/user/:id", handlers.DelUeserHandler)

	router.Run(":" + envs.ServerEnvs.BACKEND_PORT)
}
