package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/domain"
	"github.com/privatesquare/bkst-oauth-api/services"
	"net/http"
)

func NewUsersHandler(s services.UsersService) UsersHandler {
	return &usersHandler{Service: s}
}

type UsersHandler interface {
	Login(ctx *gin.Context)
}

type usersHandler struct {
	Service services.UsersService
}

func (uh *usersHandler) Login(ctx *gin.Context) {
	usersService := services.NewUsersService(uh.Service)

	login, err := uh.parseLogin(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	user, restErr := usersService.Login(*login)
	if restErr != nil {
		ctx.JSON(restErr.Status, restErr)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uh *usersHandler) parseLogin(ctx *gin.Context) (*domain.Login, *errors.RestErr) {
	login := new(domain.Login)
	if err := ctx.ShouldBindJSON(login); err != nil {
		logger.Info(invalidPayloadMsg)
		return nil, errors.BadRequestError(invalidPayloadMsg)
	}
	return login, nil
}
