package handlers

import (
	"errors"
	skiplist "golabs/src/services/skiplist"

	"github.com/gin-gonic/gin"
)

type SkipListHandler struct {
	skiplistService skiplist.SkipListService
}

func NewSkiplistHandler() *SkipListHandler {
	return &SkipListHandler{
		skiplistService: skiplist.NewSkipList(),
	}
}

func (handler *SkipListHandler) Seed(c *gin.Context) {
	size := handler.skiplistService.Seed()

	c.JSON(200, gin.H{
		"status": "seeded",
		"data": gin.H{
			"size": size,
		},
	})
}

func (handler *SkipListHandler) Delete(c *gin.Context) {
	var request GetNodeKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	deletedValue, err := handler.skiplistService.Delete(request.Key)

	if errors.Is(err, skiplist.ErrNotFound) {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node deleted",
		"data": gin.H{
			"deletedValue": deletedValue,
		},
	})
}

func (handler *SkipListHandler) Insert(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	oldValue, replaced := handler.skiplistService.Insert(request.Key, request.Value)

	c.JSON(201, gin.H{
		"status": "node added",
		"data": gin.H{
			"oldValue": oldValue,
			"replaced": replaced,
		},
	})
}

func (handler *SkipListHandler) Search(c *gin.Context) {
	var request GetNodeKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	valueFound, err := handler.skiplistService.Search(request.Key)

	if errors.Is(err, skiplist.ErrNotFound) {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node found",
		"data": gin.H{
			"valueFound": valueFound,
		},
	})
}

func (handler *SkipListHandler) Contains(c *gin.Context) {
	var request GetNodeKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	found, err := handler.skiplistService.Contains(request.Key)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node found",
		"data": gin.H{
			"found": found,
		},
	})
}
