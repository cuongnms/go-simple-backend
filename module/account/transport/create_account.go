package transport

import (
	"github.com/gin-gonic/gin"
	"go-simple-backend/common"
	"go-simple-backend/module/account/biz"
	"go-simple-backend/module/account/model"
	storage "go-simple-backend/module/account/storage"
	"net/http"
)

func CreateAccount(appCtx common.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		var data model.Account
		db := appCtx.GetDatabaseConnection()
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}
		storage := storage.NewSqlStorage(db)
		biz := biz.CreateAccountStorage(storage)
		if err := biz.CreateAccount(context.Request.Context(), &data); err != nil {
			context.JSON(http.StatusBadRequest, err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
