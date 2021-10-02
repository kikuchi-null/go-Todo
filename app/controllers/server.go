package controllers

import (
	"fmt"
	"log"
	"tasks/app/models"
	"tasks/app/pkg/config"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-ini/ini.v1"
)

type PageList struct {
	Top     string
	Signup  string
	Login   string
	Index   string
	Profile string
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

	// related to task.
	router.GET("/tasks", index)
	router.GET("/profile", profile)                // List view of tasks
	router.GET("/tasks/create", create)            // Create
	router.POST("/tasks/save", save)               // Save
	router.GET("/tasks/edit/:task_id", edit)       // Edit
	router.POST("/tasks/update/:task_id", update)  // Update
	router.POST("/tasks/delete/:task_id", delete)  // Delete
	router.GET("/tasks/confirm/:task_id", confirm) // Confirm

	router.Run(config.Config.Port)
}

func LoadPageList() PageList {
	cfg, err := ini.Load("./controllers/page.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Pages := PageList{
		Login:   cfg.Section("auth").Key("login").String(),
		Signup:  cfg.Section("auth").Key("signup").String(),
		Index:   cfg.Section("main").Key("index").String(),
		Profile: cfg.Section("main").Key("profile").String(),
		Top:     cfg.Section("main").Key("top").String(),
		Create:  cfg.Section("task").Key("create").String(),
		Confirm: cfg.Section("task").Key("confirm").String(),
		Edit:    cfg.Section("task").Key("edit").String(),
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
