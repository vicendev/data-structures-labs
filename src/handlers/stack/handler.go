package handlers

import (
	"github.com/gin-gonic/gin"

	"golabs/src/services/stack"
)

type StackHandler struct {
	stackService stack.StackService
}

func NewStackHandler() *StackHandler {
	return &StackHandler{
		stackService: stack.NewStackService(),
	}
}

func (handler *StackHandler) Initialize(c *gin.Context) {
	var request InitializeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	handler.stackService.Initialize(request.Capacity)

	c.JSON(200, gin.H{
		"status":   "initialized",
		"capacity": request.Capacity,
	})

}

func (handler *StackHandler) Push(c *gin.Context) {
	var request PushRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	err := handler.stackService.Push(request.Value)

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"value":  request.Value,
	})
}

func (handler *StackHandler) Pop(c *gin.Context) {
	response, err := handler.stackService.Pop()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"value":  response,
	})
}

func (handler *StackHandler) Peek(c *gin.Context) {
	response, err := handler.stackService.Peek()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"value":  response,
	})
}

func (handler *StackHandler) Size(c *gin.Context) {
	response, err := handler.stackService.Size()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": "Ok",
		"size":   response,
	})
}

func (handler *StackHandler) IsEmpty(c *gin.Context) {
	response, err := handler.stackService.IsEmpty()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"empty":  response,
	})
}
