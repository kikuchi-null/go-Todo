package main

import (
	"fmt"
	"tasks/app/models"
)

func main() {

	// gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.TestMode)

	// controllers.StartServer()

	u, _ := models.GetUserByID(1)
	fmt.Println(u)
	sess, _ := u.CreateSession()
	fmt.Println(sess)
}
