package notemodel

import (
	"errors"
	"note_server/common"
)

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

var (
	ErrTitleCannotBeBlank = errors.New("title can not be blank")
	ErrItemNotFound       = errors.New("item not found")
)
