package controllers

import (
	"log"
	"net/http"
	"tasks/app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	_, err := session(c)

	if err != nil {

		if c.Request.Method == http.MethodGet {
			c.HTML(http.StatusOK, LoadPageList().Signup, gin.H{})

		} else if c.Request.Method == http.MethodPost {
			u := models.User{
				UUID:      models.CreateUUID().String(),
				Name:      c.PostForm("name"),
				Email:     c.PostForm("email"),
				Password:  models.Encrypt(c.PostForm("password")),
				Create_At: time.Now(),
				Update_At: time.Now(),
			}
			err := u.CreateUser()
			if err != nil {
				log.Println(err)
				c.HTML(http.StatusOK, LoadPageList().Signup, gin.H{
					"message": err,
				})
			}

			c.Redirect(http.StatusFound, "/login")

		}

	} else {

		c.Redirect(http.StatusFound, "/tasks")
	}
}

func login(c *gin.Context) {
	_, err := session(c)

	if err != nil {
		c.HTML(http.StatusOK, LoadPageList().Login, gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/tasks")
	}
}

func authenticate(c *gin.Context) {
	user, err := models.GetUserByEmail(c.PostForm("email"))
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/login")
	}

	if user.Password == models.Encrypt(c.PostForm("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		c.SetCookie("gin_cookie", session.UUID, 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusFound, "/tasks")

	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func logout(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{
			UUID: cookie,
		}

		session.DeleteSessionByUUID()
	}

	c.Redirect(http.StatusFound, "/login")

}
