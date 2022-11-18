package notebusiness

import (
	"context"
	"errors"
	notemodel "note_server/module/noteitem/model"
)

type UpdateNoteItemStorage interface {
	FindNote(
		ctx context.Context,
		condition map[string]interface{},
	) (*notemodel.NoteItem, error)

	UpdateNote(
		ctx context.Context,
		condition map[string]interface{},
		data *notemodel.NoteItem,
	) error
}

type updateBusiness struct {
	store UpdateNoteItemStorage
}

func NewUpdateNoteItemBusiness(store UpdateNoteItemStorage) *updateBusiness {
	return &updateBusiness{store: store}
}

func (biz *updateBusiness) UpdateNote(ctx context.Context,
	condition map[string]interface{},
	data *notemodel.NoteItem,
) error {
	oldData, err := biz.store.FindNote(ctx, condition)

	if err != nil {
		return err
	}

	if oldData.Status == "Deleted" {
		return errors.New("data has been deleted")
	}

	if err := biz.store.UpdateNote(ctx, condition, data); err != nil {
		return err
	}

	return nil
}
