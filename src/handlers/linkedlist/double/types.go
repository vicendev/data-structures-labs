package handlers

import service "golabs/src/services/linkedlist/double"

type NodeValue struct {
	Value string `json:"value" binding:"required"`
}

type GetNodeIndex struct {
	Index *int `form:"index" binding:"required,gte=0"`
}

type NodeInsert struct {
	SearchValue string `json:"searchValue" binding:"required"`
	Value       string `json:"value" binding:"required"`
}

type NodeInsertAt struct {
	Index *int   `json:"index" binding:"required,gte=0"`
	Value string `json:"value" binding:"required"`
}

type NodeView struct {
	Value string  `json:"data"`
	Prev  *string `json:"prev,omitempty"`
	Next  *string `json:"next,omitempty"`
}

func ToNodeView(node *service.Node) *NodeView {
	if node == nil {
		return nil
	}
	var prev, next *string
	if node.Prev != nil {
		prev = &node.Prev.Value
	}
	if node.Next != nil {
		next = &node.Next.Value
	}
	return &NodeView{
		Value: node.Value,
		Prev:  prev,
		Next:  next,
	}
}
