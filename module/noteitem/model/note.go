package notemodel

import (
	"fmt"
	"note_server/common"
	"strings"
)

const EntityName = "Note"

type NoteItem struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
	Content         string `json:"content" gorm:"column:content"`
	Category        string `json:"category" gorm:"column:category"`
	Status          string `json:"status" gorm:"column:status"`
}

func (NoteItem) TableName() string {
	return "note_item"
}

type NoteItemCreate struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title"`
	Content         string `json:"content" gorm:"column:content"`
	Category        string `json:"category" gorm:"column:category"`
	Status          string `json:"status" gorm:"column:status"`
}

func (data *NoteItemCreate) Validate() error {
	dataNames := map[string]string{
		"title":    data.Title,
		"content":  data.Content,
		"category": data.Category,
	}

	for k, v := range dataNames {
		v = strings.TrimSpace(v)

		if v == "" {
			return ErrorFieldIsEmpty(k)
		}
	}

	return nil
}

func (NoteItemCreate) TableName() string {
	return NoteItem{}.TableName()
}

type NoteItemUpdate struct {
	Title    *string `json:"title" gorm:"column:title"`
	Content  *string `json:"content" gorm:"column:content"`
	Category *string `json:"category" gorm:"column:category"`
	Status   *string `json:"status" gorm:"column:status"`
}

func (NoteItemUpdate) TableName() string {
	return NoteItem{}.TableName()
}

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}
