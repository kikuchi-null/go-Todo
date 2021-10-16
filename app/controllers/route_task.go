package controllers

import (
	"log"
	"net/http"
	"strconv"
	"tasks/app/models"

	"github.com/gin-gonic/gin"
)

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
		c.Redirect(http.StatusFound, "/tasks")
	} else {
		if c.Request.Method == http.MethodPost {
			user, err := session.GetUserBySession()
			if err != nil {
				log.Println(err)
			}

			user.CreateTask(c.PostForm("content"))
			c.Redirect(http.StatusFound, "/tasks")
		}
	}
	c.Redirect(http.StatusFound, "/tasks")
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

		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			log.Println(err)
		}

		task, err := models.GetTask(task_id)
		if err != nil {
			log.Println(err)
		}

		if c.Request.Method == http.MethodGet {

			c.HTML(http.StatusOK, LoadPageList().Edit, gin.H{
				"Content": task.Content,
				"Task_ID": task.Task_ID,
			})
		}

	}
	c.Redirect(http.StatusFound, "/tasks")
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

		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			log.Println(err)
		}
		task, err := models.GetTask(task_id)
		if err != nil {
			log.Println(err)
		}

		task.UpdateTask(c.PostForm("content"))
		c.Redirect(http.StatusFound, "/tasks")
	}

}

func confirm(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			log.Println(err)
		}

		task, err := models.GetTask(task_id)
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().Confirm, gin.H{
			"Task_ID":   task.Task_ID,
			"Create_At": task.Create_At,
			"Content":   task.Content,
			"Update_At": task.Update_At,
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

		task_id, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			log.Println(err)
		}

		task, err := models.GetTask(task_id)
		if err != nil {
			log.Println(err)
		}

		task.DeleteTask()
		c.Redirect(http.StatusFound, "/tasks")

	}
}
