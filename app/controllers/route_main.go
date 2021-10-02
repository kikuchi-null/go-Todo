package controllers

import (
	"log"
	"net/http"
	"strconv"
	"tasks/app/models"

	"github.com/gin-gonic/gin"
)

func top(c *gin.Context) {

	c.HTML(http.StatusOK, LoadPageList().Top, gin.H{
		"message": "the simplest Task management application with Golang.",
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

		Tasks, err := user.GetTasksByUser()
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().Index, gin.H{
			"Name":  user.Name,
			"Tasks": Tasks,
		})
	}
}

func create(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		c.HTML(http.StatusOK, LoadPageList().Create, gin.H{})
	}

}

func save(c *gin.Context) {
	session, err := session(c)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/Tasks")
	} else {
		if c.Request.Method == http.MethodPost {
			user, err := session.GetUserBySession()
			if err != nil {
				log.Println(err)
			}

			user.CreateTask(c.PostForm("content"))
			c.Redirect(http.StatusFound, "/Tasks")
		}
	}
	c.Redirect(http.StatusFound, "/Tasks")
}

func edit(c *gin.Context) {
	session, err := session(c)

	if err != nil {
		c.Redirect(http.StatusFound, "/login")

	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		Task_id, err := strconv.Atoi(c.Param("Task_id"))
		if err != nil {
			log.Println(err)
		}

		Task, err := models.GetTask(Task_id)
		if err != nil {
			log.Println(err)
		}

		if c.Request.Method == http.MethodGet {

			c.HTML(http.StatusOK, LoadPageList().Edit, gin.H{
				"Content": Task.Content,
				"Task_ID": Task.Task_ID,
			})
		}

	}
	c.Redirect(http.StatusFound, "/Tasks")
}

func update(c *gin.Context) {
	session, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")

	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		Task_id, err := strconv.Atoi(c.Param("Task_id"))
		if err != nil {
			log.Println(err)
		}
		Task, err := models.GetTask(Task_id)
		if err != nil {
			log.Println(err)
		}

		Task.UpdateTask(c.PostForm("content"))
		c.Redirect(http.StatusFound, "/Tasks")
	}

}

func confirm(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		Task_id, err := strconv.Atoi(c.Param("Task_id"))
		if err != nil {
			log.Println(err)
		}

		Task, err := models.GetTask(Task_id)
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().Confirm, gin.H{
			"Task_ID":   Task.Task_ID,
			"Create_At": Task.Create_At,
			"Content":   Task.Content,
			"Update_At": Task.Update_At,
		})
	}

}

func delete(c *gin.Context) {
	session, err := session(c)

	if err != nil {
		c.Redirect(http.StatusFound, "/login")

	} else {
		_, err := session.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		Task_id, err := strconv.Atoi(c.Param("Task_id"))
		if err != nil {
			log.Println(err)
		}

		Task, err := models.GetTask(Task_id)
		if err != nil {
			log.Println(err)
		}

		Task.DeleteTask()
		c.Redirect(http.StatusFound, "/Tasks")

	}
}
