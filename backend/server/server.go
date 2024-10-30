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
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
		database.DB.AutoMigrate(&models.User{})
	}

}

func StartServer() {
	InitRotes()
}
