package handler

import (
	"net/http"
	"os"

	"github.com/chris10ml/go-web-products/pkg/web"
	"github.com/gin-gonic/gin"
)

func AuthTokenHandler(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Enter token"))
		return
	}

	if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Invalid token"))
		return
	}

	ctx.Next()
}
