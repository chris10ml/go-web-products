package handler

import (
	"os"

	"github.com/chris10ml/go-web-products/pkg/web"
	"github.com/gin-gonic/gin"
)

func AuthTokenHandler(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")

	if token == "" {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Enter token"))
		return
	}

	if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Invalid token"))
		return
	}

	ctx.Next()
}
