package notebusiness

import (
	"context"
	notemodel "note_server/module/noteitem/model"
)

type CreateNoteItemStorage interface {
	CreateNewNote(ctx context.Context, data *notemodel.NoteItemCreate) error
}

type createBusiness struct {
	store CreateNoteItemStorage
}

func NewCreateNoteItemBusiness(store CreateNoteItemStorage) *createBusiness {
	return &createBusiness{store: store}
}

func (biz *createBusiness) CreateNewNote(ctx context.Context, data *notemodel.NoteItemCreate) error {
	data.Status = "Pending"

	if err := biz.store.CreateNewNote(ctx, data); err != nil {
		return err
	}

	return nil
}
