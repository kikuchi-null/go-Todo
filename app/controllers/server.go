package controllers

import (
	"fmt"
	"log"
	"todo/app/models"
	"todo/app/pkg/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-ini/ini.v1"
)

type PageList struct {
	Top     string
	Signup  string
	Login   string
	Index   string
	Create  string
	Edit    string
	Confirm string
}

func init() {
	LoadPageList()
}

func StartServer() {
	router := gin.Default()
	router.LoadHTMLGlob(config.Config.Templates)
	router.Static("/static/", "./views/")

	// Related to Auth.
	router.GET("/", top)                       // Top Page
	router.GET("/signup", signup)              // Signup Page
	router.POST("/signup", signup)             // Signup Page
	router.GET("/login", login)                // Login
	router.POST("/authenticate", authenticate) // Authenticate
	router.GET("/logout", logout)              // Logout

	// related to Todo.
	router.GET("/todos", index)                    // List view of Todos
	router.GET("/todos/create", create)            // Create
	router.POST("/todos/save", save)               // Save
	router.GET("/todos/edit/:todo_id", edit)       // Edit
	router.POST("/todos/update/:todo_id", update)  // Update
	router.GET("/todos/confirm/:todo_id", confirm) // Confirm
	router.POST("/todos/delete/:todo_id", delete)  // Delete

	router.Run(config.Config.Port)
}

func LoadPageList() PageList {
	cfg, err := ini.Load("./controllers/page.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Pages := PageList{
		Top:     cfg.Section("page").Key("top").String(),
		Signup:  cfg.Section("page").Key("signup").String(),
		Login:   cfg.Section("page").Key("login").String(),
		Index:   cfg.Section("page").Key("index").String(),
		Create:  cfg.Section("page").Key("create").String(),
		Edit:    cfg.Section("page").Key("edit").String(),
		Confirm: cfg.Section("page").Key("confirm").String(),
	}

	return Pages
}

func session(c *gin.Context) (session models.Session, err error) {
	cookie, err := c.Cookie("gin_cookie")
	if err == nil {
		session = models.Session{
			UUID: cookie,
		}

		valid, _ := session.IsSession()
		if !valid {
			err = fmt.Errorf("invalid session")
		}
	}

	return session, err
}
