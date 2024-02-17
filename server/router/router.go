package router

import (
	handler "getting-statistics-mirea/server/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(hand *handler.Handler) *gin.Engine {
	router := gin.New()
	router.Static("/client", "client")
	router.POST("/", hand.GetResult)
	router.GET("/", func(c *gin.Context) {
		c.File("client/index.html")
	})

	return router
}

func Start(router *gin.Engine, addr string) error {
	return router.Run(addr)
}
