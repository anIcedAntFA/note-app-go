package notetransport

import (
	"net/http"
	common "note_server/common"
	"note_server/components/appctx"
	notebusiness "note_server/module/noteitem/business"
	notemodel "note_server/module/noteitem/model"
	notestorage "note_server/module/noteitem/storage"

	"github.com/gin-gonic/gin"
)

func HandleListNotes(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := ctx.ShouldBind(&pagingData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		pagingData.Fulfill()

		var filter notemodel.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result []notemodel.NoteItem

		storage := notestorage.NewMySQLStorage(db)
		business := notebusiness.ListNoteItemsStorage(storage)

		result, err := business.ListNotes(ctx.Request.Context(), &filter, &pagingData)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
