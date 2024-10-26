package main

import "backend/server"

func init() {
	server.InitServer()
}

func main() {
	server.StartServer()
}

// func main() {
// 	r := gin.Default()
// 	r.Use(cors.Default())

// 	errEnvs := envs.LoadEnvs()
// 	if errEnvs != nil {
// 		log.Fatal("Ошибка инициализации ENV: ", errEnvs)
// 	} else {
// 		log.Println("Инициализация ENV прошла успешно")
// 	}

// 	r.GET("/", handlers.GetHandler)

// 	r.Run(":" + envs.ServerEnvs.BACKEND_PORT)
// }
