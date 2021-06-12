package users_api

import (
	"github.com/go-resty/resty/v2"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/domain"
	"net/http"
)

var (
	baseUrl string
)

const (
	usersLoginApiPath = "/users/login"
)

type UsersApiCfg struct {
	Url string `mapstructure:"USERS_API_URL"`
}

func (cfg *UsersApiCfg) SetConfig() {
	baseUrl = cfg.Url
}

func NewUsersStore(client *resty.Client) UsersStore {
	return &usersStore{
		Client: client,
	}
}

type UsersStore interface {
	Login(domain.Login) (*domain.User, *errors.RestErr)
}

type usersStore struct {
	Client *resty.Client
}

func (us *usersStore) Login(login domain.Login) (*domain.User, *errors.RestErr) {
	user := new(domain.User)
	restErr := new(errors.RestErr)
	us.Client.SetHostURL(baseUrl)
	resp, err := us.Client.R().SetBody(login).SetResult(user).SetError(restErr).Post(usersLoginApiPath)
	if err != nil {
		logger.Error(err.Error(), err)
		return nil, errors.InternalServerError()
	}
	if resp.StatusCode() == http.StatusOK {
		return user, nil
	} else {
		return nil, restErr
	}
}
