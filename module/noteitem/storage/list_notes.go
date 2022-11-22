package notestorage

import (
	"context"
	"note_server/common"
	notemodel "note_server/module/noteitem/model"
)

func (s *mysqlStorage) ListNotes(
	ctx context.Context,
	filter *notemodel.Filter,
	paging *common.Paging,
) ([]notemodel.NoteItem, error) {
	var result []notemodel.NoteItem

	db := s.db.Table(notemodel.NoteItem{}.TableName()).
		Where("status IN ?", []string{"Pending", "Completed"})

	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}

		if len(f.Category) > 0 {
			db = db.Where("category in (?)", f.Category)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Offset(offset).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return result, nil
}
