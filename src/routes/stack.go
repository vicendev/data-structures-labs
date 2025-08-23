package routes

import (
	handlers "golabs/src/handlers/stack"

	"github.com/gin-gonic/gin"
)

func RegisterStackRoutes(r *gin.Engine) {

	h := handlers.NewStackHandler()

	g := r.Group("/stack")
	{
		g.POST("/initialize", h.Initialize)
		g.POST("/push", h.Push)
		g.GET("/pop", h.Pop)
		g.GET("/peek", h.Peek)
		g.GET("/size", h.Size)
		g.GET("/is-empty", h.IsEmpty)
	}
}
