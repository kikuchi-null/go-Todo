package controllers

import (
	"log"
	"net/http"
	"strconv"
	"todo/app/models"

	"github.com/gin-gonic/gin"
)

func top(c *gin.Context) {

	c.HTML(http.StatusOK, LoadPageList().Top, gin.H{
		"message": "the simplest Task management application with Golang.",
	})
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

		todos, err := user.GetTodosByUser()
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().Index, gin.H{
			"Name":  user.Name,
			"Todos": todos,
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
		c.Redirect(http.StatusFound, "/todos")
	} else {
		if c.Request.Method == http.MethodPost {
			user, err := session.GetUserBySession()
			if err != nil {
				log.Println(err)
			}

			user.CreateTodo(c.PostForm("content"))
			c.Redirect(http.StatusFound, "/todos")
		}
	}
	c.Redirect(http.StatusFound, "/todos")
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

		todo_id, err := strconv.Atoi(c.Param("todo_id"))
		if err != nil {
			log.Println(err)
		}

		todo, err := models.GetTodo(todo_id)
		if err != nil {
			log.Println(err)
		}

		if c.Request.Method == http.MethodGet {

			c.HTML(http.StatusOK, LoadPageList().Edit, gin.H{
				"Content": todo.Content,
				"Todo_ID": todo.Todo_ID,
			})
		}

	}
	c.Redirect(http.StatusFound, "/todos")
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

		todo_id, err := strconv.Atoi(c.Param("todo_id"))
		if err != nil {
			log.Println(err)
		}
		todo, err := models.GetTodo(todo_id)
		if err != nil {
			log.Println(err)
		}

		todo.UpdateTodo(c.PostForm("content"))
		c.Redirect(http.StatusFound, "/todos")
	}

}

func confirm(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		todo_id, err := strconv.Atoi(c.Param("todo_id"))
		if err != nil {
			log.Println(err)
		}

		todo, err := models.GetTodo(todo_id)
		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, LoadPageList().Confirm, gin.H{
			"Todo_ID":   todo.Todo_ID,
			"Create_At": todo.Create_At,
			"Content":   todo.Content,
			"Update_At": todo.Update_At,
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

		todo_id, err := strconv.Atoi(c.Param("todo_id"))
		if err != nil {
			log.Println(err)
		}

		todo, err := models.GetTodo(todo_id)
		if err != nil {
			log.Println(err)
		}

		todo.DeleteTodo()
		c.Redirect(http.StatusFound, "/todos")

	}
}
