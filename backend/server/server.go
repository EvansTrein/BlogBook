package server

import (
	"backend/envs"
	"log"
)

func InitServer() {
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
	} else {
		log.Println("Инициализация ENV прошла успешно")
	}
}

func StartServer() {
	InitRotes()
}
