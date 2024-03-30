package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, message string, data any, status int, code int) {
	ctx.JSON(status, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func FialWithResponse(ctx *gin.Context, err error) {
	if customErr, ok := err.(*Exception); ok {
		resultBody := gin.H{
			"code":    customErr.Code(),
			"message": customErr.Error(),
		}
		if len(customErr.ErrorMessage()) != 0 {
			resultBody["errors"] = customErr.ErrorMessage()
		}
		ctx.JSON(customErr.Status(), resultBody)
		return
	}
	FialWithResponse(ctx, ErrInternalServer)
}

func OkWithResponse(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}
