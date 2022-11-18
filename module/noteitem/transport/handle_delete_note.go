package notetransport

import (
	"net/http"
	"note_server/common"
	"note_server/components/appctx"
	notebusiness "note_server/module/noteitem/business"
	notestorage "note_server/module/noteitem/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleDeleteNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := notestorage.NewMySQLStorage(db)
		business := notebusiness.NewDeleteNoteItemBusiness(storage)

		if err := business.DeleteNote(ctx.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
