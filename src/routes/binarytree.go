package routes

import (
	handlers "golabs/src/handlers/tree/binarytree"

	"github.com/gin-gonic/gin"
)

func RegisterBinaryTreeRoutes(r *gin.Engine) {

	h := handlers.NewBinaryTreeHandler()

	g := r.Group("/binarytree")
	{
		g.POST("/upsert", h.Upsert)
		g.GET("/search", h.Search)
		g.GET("/seed", h.Seed)
		g.DELETE("/delete", h.Delete)
		g.GET("debug-tree", h.DebugTree)
		g.GET("/reset", h.Reset)
	}
}
