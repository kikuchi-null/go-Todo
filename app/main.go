package main

import (
	"tasks/app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	controllers.StartServer()

}
