package hashtable

import (
	"errors"
	hashtable "golabs/src/services/hashtable"
	"golabs/src/services/skiplist"

	"github.com/gin-gonic/gin"
)

type HashTableHandler struct {
	hashtableService hashtable.HashTableService
}

func NewHashTableHandler() *HashTableHandler {
	return &HashTableHandler{
		hashtableService: hashtable.NewHashTable(),
	}
}

func (handler *HashTableHandler) Seed(c *gin.Context) {
	var request SeedValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	size := handler.hashtableService.Seed()

	c.JSON(200, gin.H{
		"status": "seeded",
		"data": gin.H{
			"size": size,
		},
	})
}

func (handler *HashTableHandler) Reset(c *gin.Context) {
	handler.hashtableService.Reset()
	c.JSON(200, gin.H{
		"status": "hash table has been reseted",
	})
}

func (handler *HashTableHandler) Clear(c *gin.Context) {
	handler.hashtableService.Clear()
	c.JSON(200, gin.H{
		"status": "hash table has been cleared",
	})
}

func (handler *HashTableHandler) Upsert(c *gin.Context) {
	var request PairValue

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	oldValue, replaced, err := handler.hashtableService.Upsert(request.Key, request.Value, hashtable.HashFnType(request.HashFnType))

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"status": "node added",
		"data": gin.H{
			"oldValue": oldValue,
			"replaced": replaced,
		},
	})

}

func (handler *HashTableHandler) Get(c *gin.Context) {
	var request GetPairKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	valueFound, err := handler.hashtableService.Get(request.Key)

	if errors.Is(err, skiplist.ErrNotFound) {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "pair found",
		"data": gin.H{
			"valueFound": valueFound,
		},
	})
}

func (handler *HashTableHandler) Delete(c *gin.Context) {
	var request GetPairKey

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Query format", "details": err.Error()})
		return
	}

	deletedValue, err := handler.hashtableService.Delete(request.Key)

	if errors.Is(err, skiplist.ErrNotFound) {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "pair deleted",
		"data": gin.H{
			"deletedValue": deletedValue,
		},
	})
}
