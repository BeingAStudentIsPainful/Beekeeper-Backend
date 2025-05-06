package types

type EntryType string

const (
	EntryTypeLog  EntryType = "log"
	EntryTypeTask EntryType = "task"
)

type CreateEntryInput struct {
	Content  string    `json:"content" binding:"required"`
	Type     EntryType `json:"type" binding:"required,oneof=log task"`
	HiveName string    `json:"hiveName" binding:"required"`
}
