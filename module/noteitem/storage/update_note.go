package notestorage

import (
	"context"
	notemodel "note_server/module/noteitem/model"
)

func (s *mysqlStorage) UpdateNote(
	ctx context.Context,
	condition map[string]interface{},
	data *notemodel.NoteItem,
) error {
	if err := s.db.Where(condition).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
