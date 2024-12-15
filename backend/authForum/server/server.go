package server

import (
	"authForum/database"
	"authForum/envs"
	"authForum/logging"
)

func InitServer() {
	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		logging.LogError.Fatal("ENV initialization error: ", errEnvs)
	} else {
		logging.LogDebug.Println("ENV initialization was successful")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		logging.LogError.Fatal("Database connection error: ", errDatabase)
	} else {
		logging.LogDebug.Println("Successful connection to the database")
	}
}

func StartServer() {
	InitRotes()
}
