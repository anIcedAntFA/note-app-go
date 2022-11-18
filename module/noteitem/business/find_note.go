package notebusiness

import (
	"context"
	notemodel "note_server/module/noteitem/model"
)

type FindNoteItemStorage interface {
	FindNote(
		ctx context.Context,
		condition map[string]interface{},
	) (*notemodel.NoteItem, error)
}

type findBusiness struct {
	store FindNoteItemStorage
}

func NewFindNoteItemBusiness(store FindNoteItemStorage) *findBusiness {
	return &findBusiness{store: store}
}

func (biz *findBusiness) FindNote(
	ctx context.Context,
	condition map[string]interface{},
) (*notemodel.NoteItem, error) {
	data, err := biz.store.FindNote(ctx, condition)

	if err != nil {
		return nil, err
	}

	return data, nil
}
