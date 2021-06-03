package cassandra

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/domain/access_token"
)

func NewAccessTokenStore() AccessTokenStore {
	return &accessTokenStore{}
}

type AccessTokenStore interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type accessTokenStore struct{}

func (r *accessTokenStore) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := NewSession()
	if err != nil {
		logger.Error("Unable to create db session", err)
		return nil, errors.InternalServerError()
	}
	defer session.Close()
	return nil, errors.InternalServerError()
}
