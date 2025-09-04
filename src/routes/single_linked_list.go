package routes

import (
	handlers "golabs/src/handlers/linkedlist/single"

	"github.com/gin-gonic/gin"
)

func RegisterSingleLinkedListRoutes(r *gin.Engine) {

	h := handlers.NewSingleLinkedListHandler()

	g := r.Group("/single-linked-list")
	{
		g.POST("/add-first", h.AddFirst)
		g.POST("/add-last", h.AddLast)
		g.GET("/clear", h.Clear)
		g.POST("/find", h.Find)
		g.POST("/get-at", h.GetAt)
		g.POST("/insert-at", h.InsertAt)
		g.POST("/insert-after", h.InsertAfter)
		g.DELETE("/remove", h.Remove)
		g.DELETE("/remove-at", h.RemoveAt)
		g.DELETE("/remove-first", h.RemoveFirst)
		g.DELETE("/remove-last", h.RemoveLast)
	}
}
