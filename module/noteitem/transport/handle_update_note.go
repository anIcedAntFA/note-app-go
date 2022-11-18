package notetransport

import (
	"net/http"
	"note_server/components/appctx"
	notebusiness "note_server/module/noteitem/business"
	notemodel "note_server/module/noteitem/model"
	notestorage "note_server/module/noteitem/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleUpdateNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data notemodel.NoteItem

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := notestorage.NewMySQLStorage(db)
		business := notebusiness.NewUpdateNoteItemBusiness(storage)

		if err := business.UpdateNote(
			ctx.Request.Context(),
			map[string]interface{}{"id": id},
			&data,
		); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
