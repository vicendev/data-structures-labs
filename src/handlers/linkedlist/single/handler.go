package handlers

import (
	linkedlist "golabs/src/services/linkedlist/single"

	"github.com/gin-gonic/gin"
)

type SingleLinkedListHandler struct {
	singleLinkedListService linkedlist.SingleLinkedListService
}

func NewSingleLinkedListHandler() *SingleLinkedListHandler {
	return &SingleLinkedListHandler{
		singleLinkedListService: linkedlist.NewSingleLinkedList(),
	}
}

func (handler *SingleLinkedListHandler) AddFirst(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	handler.singleLinkedListService.AddFirst(request.Value)

	c.JSON(200, gin.H{
		"status": "node added to head",
		"value":  request.Value,
	})

}

func (handler *SingleLinkedListHandler) AddLast(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	handler.singleLinkedListService.AddLast(request.Value)

	c.JSON(200, gin.H{
		"status": "node added to head",
		"value":  request.Value,
	})

}

func (handler *SingleLinkedListHandler) Clear(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "single linked list cleared",
	})
}

func (handler *SingleLinkedListHandler) Find(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	node, err := handler.singleLinkedListService.Find(request.Value)

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
		"value":  node,
	})

}

func (handler *SingleLinkedListHandler) GetAt(c *gin.Context) {
	var request NodeIndex

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	response, err := handler.singleLinkedListService.GetAt(*request.Index)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		c.JSON(404, gin.H{"error": "node not found"})
		return
	}

	c.JSON(200, gin.H{
		"status": "node found",
		"value":  response,
	})

}

func (handler *SingleLinkedListHandler) InsertAt(c *gin.Context) {
	var request NodeInsertAt

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	if err := handler.singleLinkedListService.InsertAt(*request.Index, request.Value); err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node inserted successfully",
		"value":  request,
	})

}

func (handler *SingleLinkedListHandler) InsertAfter(c *gin.Context) {
	var request NodeInsertAfter

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	if err := handler.singleLinkedListService.InsertAfter(request.SearchValue, request.Value); err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node inserted after successfully",
		"value":  request,
	})

}

func (handler *SingleLinkedListHandler) Remove(c *gin.Context) {
	var request NodeValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	response, err := handler.singleLinkedListService.Remove(request.Value)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  response,
	})

}

func (handler *SingleLinkedListHandler) RemoveAt(c *gin.Context) {
	var request NodeIndex

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "datails": err.Error()})
		return
	}

	response, err := handler.singleLinkedListService.RemoveAt(*request.Index)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  response,
	})

}

func (handler *SingleLinkedListHandler) RemoveLast(c *gin.Context) {
	response, err := handler.singleLinkedListService.RemoveLast()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  response,
	})
}

func (handler *SingleLinkedListHandler) RemoveFirst(c *gin.Context) {
	response, err := handler.singleLinkedListService.RemoveFirst()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "node removed successfully",
		"value":  response,
	})
}
