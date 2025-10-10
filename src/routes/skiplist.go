package routes

import (
	handlers "golabs/src/handlers/skiplist"

	"github.com/gin-gonic/gin"
)

func RegisterSkipListRoutes(r *gin.Engine) {

	h := handlers.NewSkiplistHandler()

	g := r.Group("/skiplist")
	{
		g.GET("/seed", h.Seed)
		g.POST("/insert", h.Insert)
		g.GET("/search", h.Search)
		g.GET("/contains", h.Contains)
		g.DELETE("/delete", h.Delete)
	}
}
