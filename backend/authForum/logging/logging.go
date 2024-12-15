package logging

import (
	"log"
	"os"
)

var (
    LogInfo  *log.Logger
    LogDebug  *log.Logger
    LogError *log.Logger
)

func InitLogger() {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	LogInfo = log.New(os.Stdout, "LOG --> [INFO]", flags)
	LogDebug = log.New(os.Stdout, "LOG --> [DEBUG]", flags)
	LogError = log.New(os.Stderr, "LOG --> [ERROR]", flags)
}