package handlers

import (
	linkedlist "golabs/src/services/linkedlist/double"

	"github.com/gin-gonic/gin"
)

type DoubleLinkedListHandler struct {
	doubleLinkedListService linkedlist.DoubleLinkedListService
}

func NewDoubleLinkedListHandler() *DoubleLinkedListHandler {
	return &DoubleLinkedListHandler{
		doubleLinkedListService: linkedlist.NewDoubleLinkedList(),
	}
}

func (handler *DoubleLinkedListHandler) AddFirst(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	handler.doubleLinkedListService.AddFirst(request.Value)

	c.JSON(200, gin.H{
		"status": "node added to head",
		"value":  request.Value,
	})
}

func (handler *DoubleLinkedListHandler) AddLast(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	handler.doubleLinkedListService.AddLast(request.Value)

	c.JSON(200, gin.H{
		"status": "node added to tail",
		"value":  request.Value,
	})
}

func (handler *DoubleLinkedListHandler) Clear(c *gin.Context) {
	handler.doubleLinkedListService.Clear()

	c.JSON(200, gin.H{"status": "list cleared"})
}

func (handler *DoubleLinkedListHandler) Find(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	node, err := handler.doubleLinkedListService.Find(request.Value)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	if node == nil {
		c.JSON(404, gin.H{"error": "node not found"})
		return
	}

	c.JSON(200, gin.H{
		"status": "node found",
		"value":  ToNodeView(node),
	})
}

func (handler *DoubleLinkedListHandler) GetAt(c *gin.Context) {
	var params GetNodeIndex

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	node, err := handler.doubleLinkedListService.GetAt(*params.Index)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	if node == nil {
		c.JSON(404, gin.H{"error": "node not found"})
		return
	}

	c.JSON(200, gin.H{
		"status": "node found",
		"value":  ToNodeView(node),
	})

}

func (handler *DoubleLinkedListHandler) IndexOf(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	nodeIndex, err := handler.doubleLinkedListService.IndexOf(request.Value)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	if nodeIndex == -1 {
		c.JSON(404, gin.H{"error": "node not found"})
		return
	}

	c.JSON(200, gin.H{
		"status": "node found",
		"value":  nodeIndex,
	})
}

func (handler *DoubleLinkedListHandler) InsertAfter(c *gin.Context) {
	var request NodeInsert

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	if err := handler.doubleLinkedListService.InsertAfter(request.SearchValue, request.Value); err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node inserted after successfully",
		"value":  request,
	})
}

func (handler *DoubleLinkedListHandler) InsertAt(c *gin.Context) {
	var request NodeInsertAt

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	if err := handler.doubleLinkedListService.InsertAt(*request.Index, request.Value); err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node inserted at successfully",
		"value":  request,
	})
}

func (handler *DoubleLinkedListHandler) Remove(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	node, err := handler.doubleLinkedListService.Remove(request.Value)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  node,
	})
}

func (handler *DoubleLinkedListHandler) RemoveAt(c *gin.Context) {
	var params GetNodeIndex

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	node, err := handler.doubleLinkedListService.RemoveAt(*params.Index)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  node,
	})

}

func (handler *DoubleLinkedListHandler) RemoveFirst(c *gin.Context) {
	node, err := handler.doubleLinkedListService.RemoveFirst()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  node,
	})

}

func (handler *DoubleLinkedListHandler) RemoveLast(c *gin.Context) {
	node, err := handler.doubleLinkedListService.RemoveLast()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  node,
	})

}
