package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	RegisterStackRoutes(r)
	RegisterQueueRoutes(r)
	RegisterSingleLinkedListRoutes(r)
	RegisterDoubleLinkedListRoutes(r)
}
