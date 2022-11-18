package notebusiness

import (
	"context"
	"note_server/common"
	notemodel "note_server/module/noteitem/model"
)

type ListNoteItemsStorage interface {
	ListNotes(
		ctx context.Context,
		filter *notemodel.Filter,
		paging *common.Paging,
	) ([]notemodel.NoteItem, error)
}

type listBusiness struct {
	store ListNoteItemsStorage
}

func NewListNoteItemsStorage(store ListNoteItemsStorage) *listBusiness {
	return &listBusiness{store: store}
}

func (biz listBusiness) ListNotes(
	ctx context.Context,
	filter *notemodel.Filter,
	paging *common.Paging,
) ([]notemodel.NoteItem, error) {
	result, err := biz.store.ListNotes(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
