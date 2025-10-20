package handler

import (
	"errors"
	binarytree "golabs/src/services/tree/binarytree"

	"github.com/gin-gonic/gin"
)

type BinaryTreeHandler struct {
	binarytreeService binarytree.BinaryTreeService
}

func NewBinaryTreeHandler() *BinaryTreeHandler {
	return &BinaryTreeHandler{
		binarytreeService: binarytree.NewBinaryTree(),
	}
}

func (handler *BinaryTreeHandler) DebugTree(c *gin.Context) {
	handler.binarytreeService.DebugTree()

	c.JSON(201, gin.H{
		"status": "debug in console",
	})
}

func (handler *BinaryTreeHandler) Seed(c *gin.Context) {
	size := handler.binarytreeService.Seed()

	c.JSON(201, gin.H{
		"status": "seeded",
		"data": gin.H{
			"size": size,
		},
	})
}

func (handler *BinaryTreeHandler) Reset(c *gin.Context) {
	handler.binarytreeService.Reset()

	c.JSON(201, gin.H{
		"status": "binary tree restarted",
	})
}

func (handler *BinaryTreeHandler) Upsert(c *gin.Context) {
	var request NodeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	oldValue, replaced := handler.binarytreeService.Upsert(request.Key, request.Value)

	c.JSON(201, gin.H{
		"status": "node added",
		"data": gin.H{
			"oldValue": oldValue,
			"replaced": replaced,
		},
	})
}

func (handler *BinaryTreeHandler) Search(c *gin.Context) {
	var request GetNodeKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	valueFound, err := handler.binarytreeService.Search(request.Key)

	if errors.Is(err, binarytree.ErrNotFound) {
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

func (handler *BinaryTreeHandler) Delete(c *gin.Context) {
	var request GetNodeKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	deletedValue, err := handler.binarytreeService.Delete(request.Key)

	if errors.Is(err, binarytree.ErrNotFound) {
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
