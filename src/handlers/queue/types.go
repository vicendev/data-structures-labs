package handlers

type InitializeRequest struct {
	Capacity int `json:"capacity" binding:"required"`
}

type EnqueueRequest struct {
	Value string `json:"value" binding:"required"`
}
