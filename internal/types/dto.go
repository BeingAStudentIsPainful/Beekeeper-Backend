package types

type CreateEntryInput struct {
	Content string `json:"content" binding:"required"`
	HiveID  uint   `json:"hiveID" binding:"required"`
}
