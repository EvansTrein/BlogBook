package envs

import (
	"os"
	"github.com/joho/godotenv"
)

type Envs struct {
	BACKEND_PORT string
}

var ServerEnvs Envs

func LoadEnvs() error {
	
	if err := godotenv.Load(); err != nil {
		return err
	}

	ServerEnvs.BACKEND_PORT = os.Getenv("BACKEND_PORT")

	return nil
}
