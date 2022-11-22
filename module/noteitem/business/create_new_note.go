package notebusiness

import (
	"context"
	"note_server/common"
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
	if err := data.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	data.Status = "Pending"

	if err := biz.store.CreateNewNote(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(notemodel.EntityName, err)
	}

	return nil
}
