package services

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/domain"
	"github.com/privatesquare/bkst-oauth-api/interfaces/http/users_api"
)

func NewUsersService(UserStore users_api.UsersStore) UsersService {
	return &usersService{
		UserStore: UserStore,
	}
}

type UsersService interface {
	Login(domain.Login) (*domain.User, *errors.RestErr)
}

type usersService struct {
	UserStore users_api.UsersStore
}

func (s *usersService) Login(login domain.Login) (*domain.User, *errors.RestErr) {
	if err := login.Validate(); err != nil {
		logger.Info(err.Error())
		return nil, errors.BadRequestError(err.Error())
	}
	user, restErr := s.UserStore.Login(login)
	return user, restErr
}
