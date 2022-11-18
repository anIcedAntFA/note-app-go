package notestorage

import (
	"context"
	notemodel "note_server/module/noteitem/model"
)

func (s *mysqlStorage) DeleteNote(
	ctx context.Context,
	condition map[string]interface{},
) error {
	if err := s.db.Table(notemodel.NoteItem{}.TableName()).
		Where(condition).
		Updates(map[string]interface{}{"status": "deleted"}).Error; err != nil {
		return err
	}

	return nil
}
