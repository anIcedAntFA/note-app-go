package notebusiness

import (
	"context"
	"note_server/common"
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
		return common.ErrorEntityNotFound(notemodel.EntityName, err)
	}

	if oldData.Status == "Deleted" {
		return common.ErrorEntityDeleted(notemodel.EntityName, nil)
	}

	if err := biz.store.DeleteNote(ctx, condition); err != nil {
		return common.ErrorCannotDeleteEntity(notemodel.EntityName, nil)
	}

	return nil
}
