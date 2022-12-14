package notetransport

import (
	"net/http"
	common "note_server/common"
	appctx "note_server/components/appctx"
	notebusiness "note_server/module/noteitem/business"
	notemodel "note_server/module/noteitem/model"
	notestorage "note_server/module/noteitem/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewNote(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data notemodel.NoteItemCreate

		if err := ctx.ShouldBind(&data); err != nil {
			// panic(err)
			ctx.JSON(http.StatusBadRequest, common.ErrorInvalidRequest(err))

			return
		}

		//setup dependencies
		storage := notestorage.NewMySQLStorage(db)
		business := notebusiness.NewCreateNoteItemBusiness(storage)

		if err := business.CreateNewNote(ctx.Request.Context(), &data); err != nil {
			// panic(err)
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
