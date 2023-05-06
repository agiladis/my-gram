package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdentityFromCtx(ctx *gin.Context) (dataOut Claims, err error) {
	accessClaim, ok := ctx.Get("access_claim")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}

	if err = ObjectMapper(accessClaim, &dataOut); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid payload",
		})
		return
	}

	return
}
