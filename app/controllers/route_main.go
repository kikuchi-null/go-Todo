package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func top(c *gin.Context) {

	c.HTML(http.StatusOK, LoadPageList().Top, gin.H{
		"message": "The simplest task management application with Golang.",
	})
}

func profile(c *gin.Context) {
	session, err := session(c)
	if err != nil {
		c.Redirect(http.StatusOK, "/")
	} else {
		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().User, gin.H{
			"Name":      user.Name,
			"Email":     user.Email,
			"CreatedAt": user.Create_At,
		})
	}
}

func index(c *gin.Context) {
	session, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
	} else {

		user, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		tasks, err := user.GetTasksByUser()
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().Index, gin.H{
			"Name":  user.Name,
			"Tasks": tasks,
		})
	}
}
