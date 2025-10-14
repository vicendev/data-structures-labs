package hashtable

type SeedValue struct {
	HashFnType string `json:"hashFnType" binding:"required"`
}

type PairValue struct {
	Key        string `json:"key" binding:"required"`
	Value      string `json:"value" binding:"required"`
	HashFnType string `json:"hashFnType" binding:"required"`
}

type GetPairKey struct {
	Key string `form:"key" binding:"required"`
}
