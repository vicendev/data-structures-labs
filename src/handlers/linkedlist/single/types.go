package handlers

type NodeValue struct {
	Value string `json:"value" binding:"required"`
}

type NodeIndex struct {
	Index *int `json:"index" binding:"required,gte=0"`
}

type NodeInsertAt struct {
	Index *int   `json:"index" binding:"required,gte=0"`
	Value string `json:"value" binding:"required"`
}

type NodeInsertAfter struct {
	SearchValue string `json:"searchValue" binding:"required,gte=0"`
	Value       string `json:"value" binding:"required"`
}
