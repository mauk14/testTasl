package http

import (
	"github.com/gin-gonic/gin"
	"newProject/internal/service"
)

type App struct {
	greenService service.GreenService
	router       *gin.Engine
}

func setUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func NewApp(greenApi service.GreenService) *App {
	return &App{
		greenService: greenApi,
		router:       setUpRouter(),
	}
}

func (a *App) Route() *gin.Engine {
	a.router.LoadHTMLGlob("./././ui/templates/*")
	a.router.Use(recoverPanic())
	a.router.GET("/", a.getMain)
	a.router.GET("/getSettings", a.getSettings)
	a.router.GET("/getStateInstance", a.getStateInstance)
	a.router.POST("/sendMessage", a.sendMessage)
	a.router.POST("/sendFileByUrl", a.sendFileByUrl)
	return a.router
}
