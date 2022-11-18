package notestorage

import (
	"context"
	notemodel "note_server/module/noteitem/model"
)

func (s *mysqlStorage) CreateNewNote(ctx context.Context, data *notemodel.NoteItemCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
