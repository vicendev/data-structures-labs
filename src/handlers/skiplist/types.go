package handlers

type NodeValue struct {
	Key   int    `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type GetNodeKey struct {
	Key int `form:"key" binding:"required"`
}
