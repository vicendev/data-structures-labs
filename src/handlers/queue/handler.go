package handlers

import (
	"golabs/src/services/queue"

	"github.com/gin-gonic/gin"
)

type QueueHandler struct {
	queueService queue.QueueService
}

func NewQueueHandler() *QueueHandler {
	return &QueueHandler{
		queueService: queue.NewQueueService(),
	}
}

func (handler *QueueHandler) Initialize(c *gin.Context) {
	var request InitializeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	handler.queueService.Initialize(request.Capacity)

	c.JSON(200, gin.H{
		"status":   "initialized",
		"capacity": request.Capacity,
	})
}

func (handler *QueueHandler) Enqueue(c *gin.Context) {
	var request EnqueueRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	if err := handler.queueService.Enqueue(request.Value); err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Enqueued Successfully",
		"value":  request.Value,
	})
}

func (handler *QueueHandler) Dequeue(c *gin.Context) {
	response, err := handler.queueService.Dequeue()

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"value":  response,
	})
}

func (handler *QueueHandler) Tail(c *gin.Context) {
	response, err := handler.queueService.Tail()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"value":  response,
	})
}

func (handler *QueueHandler) Head(c *gin.Context) {
	response, err := handler.queueService.Head()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"value":  response,
	})
}

func (handler *QueueHandler) Size(c *gin.Context) {
	response, err := handler.queueService.Size()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": "Ok",
		"size":   response,
	})
}

func (handler *QueueHandler) IsEmpty(c *gin.Context) {
	response, err := handler.queueService.IsEmpty()

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "Ok",
		"empty":  response,
	})
}
