package server

import (
	"backend/database"
	"backend/envs"
	"backend/models"
	"log"
)

func InitServer() {
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("ENV initialization error: ", errEnvs)
	} else {
		log.Println("ENV initialization was successful")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Database connection error: ", errDatabase)
	} else {
		log.Println("Successful connection to the database")
		database.DB.AutoMigrate(&models.User{})
	}
}

func StartServer() {
	InitRotes()
}
