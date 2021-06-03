package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-oauth-api/services"
	"net/http"
	"strings"
)

func NewAccessTokenHandler(s services.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{Service: s}
}

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
}

type accessTokenHandler struct {
	Service services.AccessTokenService
}

func (ath *accessTokenHandler) GetById(ctx *gin.Context) {
	at, restErr := ath.Service.GetById(strings.TrimSpace(ctx.Param("id")))
	if restErr != nil {
		ctx.JSON(restErr.Status, restErr)
		return
	}
	ctx.JSON(http.StatusOK, at)
}
