package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/domain"
	"github.com/privatesquare/bkst-oauth-api/services"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	invalidPayloadMsg = "invalid payload"
)

func NewAccessTokenHandler(s services.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{Service: s}
}

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
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

func (ath *accessTokenHandler) Create(ctx *gin.Context) {
	at, restErr := parseAccessToken(ctx)
	if restErr != nil {
		logger.Info(restErr.Message)
		ctx.JSON(restErr.Status, restErr)
		return
	}
	restErr = ath.Service.Create(*at)
	if restErr != nil {
		logger.Info(restErr.Message)
		ctx.JSON(restErr.Status, restErr)
		return
	}
	ctx.JSON(http.StatusCreated, at)
}

func (ath *accessTokenHandler) Update(ctx *gin.Context) {
	at := new(domain.AccessToken)
	restErr := ath.Service.Update(*at)
	if restErr != nil {
		ctx.JSON(restErr.Status, restErr)
		return
	}
	ctx.JSON(http.StatusOK, at)
}

func parseAccessToken(ctx *gin.Context) (*domain.AccessToken, *errors.RestErr) {
	at := new(domain.AccessToken)
	if err := ctx.ShouldBindJSON(at); err != nil {
		logger.Info(invalidPayloadMsg, zap.NamedError("error", err))
		return nil, errors.BadRequestError(invalidPayloadMsg)
	}
	return at, nil
}
