package notestorage

import (
	"context"
	"errors"
	notemodel "note_server/module/noteitem/model"

	"gorm.io/gorm"
)

func (s *mysqlStorage) FindNote(
	ctx context.Context,
	condition map[string]interface{},
) (*notemodel.NoteItem, error) {
	var data notemodel.NoteItem

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, errors.New("Note item not found")
		}
		return nil, err
	}

	return &data, nil
}
