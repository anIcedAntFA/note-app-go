package notestorage

import (
	"context"
	"note_server/common"
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
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
