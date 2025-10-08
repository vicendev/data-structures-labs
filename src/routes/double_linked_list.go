package routes

import (
	handlers "golabs/src/handlers/linkedlist/double"

	"github.com/gin-gonic/gin"
)

func RegisterDoubleLinkedListRoutes(r *gin.Engine) {

	h := handlers.NewDoubleLinkedListHandler()

	g := r.Group("/double-linked-list")
	{
		g.POST("/add-first", h.AddFirst)
		g.POST("/add-last", h.AddLast)
		g.GET("/clear", h.Clear)
		g.POST("/find", h.Find)
		g.GET("/get-at", h.GetAt)
		g.POST("/index-of", h.IndexOf)
		g.POST("/insert-after", h.InsertAfter)
		g.POST("/insert-at", h.InsertAt)
		g.DELETE("/remove", h.Remove)
		g.DELETE("/remove-at", h.RemoveAt)
		g.DELETE("/remove-first", h.RemoveFirst)
		g.DELETE("/remove-last", h.RemoveLast)
	}
}
