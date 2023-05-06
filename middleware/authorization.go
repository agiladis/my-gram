package middleware

import (
	"my-gram/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	HeaderKey  string
	ContextKey string
)

func (h HeaderKey) String() string {
	return string(h)
}

func (h ContextKey) String() string {
	return string(h)
}

const (
	Authorization HeaderKey  = "Authorization"
	AccessClaim   ContextKey = "access_claim"
	BearerAuth    string     = "Bearer "
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(Authorization.String())
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "please login first",
			})
			return
		}

		token := strings.Split(header, BearerAuth)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token is not found",
			})
			return
		}

		var claim helper.Claims
		err := helper.VerifyToken(token[1], &claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token is not valid",
			})
			return
		}

		ctx.Set(AccessClaim.String(), claim)
		ctx.Next()
	}
}
