package types

type CreateEntryInput struct {
	Content string `json:"content" binding:"required"`
	HiveID  int    `json:"hiveID" binding:"required"`
}

type UpdateEntryInput struct {
	Content *string `json:"content"`
	HiveID  *int    `json:"hiveID"`
}
