package notebusiness

import (
	"context"
	"errors"
	notemodel "note_server/module/noteitem/model"
)

type DeleteNoteItemStorage interface {
	FindNote(
		ctx context.Context,
		condition map[string]interface{},
	) (*notemodel.NoteItem, error)

	DeleteNote(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteBusiness struct {
	store DeleteNoteItemStorage
}

func NewDeleteNoteItemBusiness(store DeleteNoteItemStorage) *deleteBusiness {
	return &deleteBusiness{store: store}
}

func (biz *deleteBusiness) DeleteNote(
	ctx context.Context,
	condition map[string]interface{},
) error {
	oldData, err := biz.store.FindNote(ctx, condition)

	if err != nil {
		return err
	}

	if oldData.Status == "Deleted" {
		return errors.New("data has been deleted")
	}

	if err := biz.store.DeleteNote(ctx, condition); err != nil {
		return err
	}

	return nil
}
