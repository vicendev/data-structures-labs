package routes

import (
	handlers "golabs/src/handlers/queue"

	"github.com/gin-gonic/gin"
)

func RegisterQueueRoutes(r *gin.Engine) {

	h := handlers.NewQueueHandler()

	g := r.Group("/queue")
	{
		g.POST("/initialize", h.Initialize)
		g.POST("/enqueue", h.Enqueue)
		g.GET("/dequeue", h.Dequeue)
		g.GET("/tail", h.Tail)
		g.GET("/head", h.Head)
		g.GET("/size", h.Size)
		g.GET("/is-empty", h.IsEmpty)
	}
}
