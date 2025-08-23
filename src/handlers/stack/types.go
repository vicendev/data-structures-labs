package handlers

type InitializeRequest struct {
	Capacity int `json:"capacity" binding:"required"`
}

type PushRequest struct {
	Value string `json:"value" binding:"required"`
}
