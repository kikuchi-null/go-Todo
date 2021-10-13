package utils

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func LoggingSettings(logFile string) {

	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logfile)

}
