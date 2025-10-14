package routes

import (
	handlers "golabs/src/handlers/hashtable"

	"github.com/gin-gonic/gin"
)

func RegisterHashTableRoutes(r *gin.Engine) {

	h := handlers.NewHashTableHandler()

	g := r.Group("/hashtable")
	{
		g.POST("/upsert", h.Upsert)
		g.POST("/seed", h.Seed)
		g.GET("/get", h.Get)
		g.DELETE("/delete", h.Delete)
	}
}
