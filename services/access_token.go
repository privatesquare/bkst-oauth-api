package services

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-oauth-api/domain/access_token"
	"github.com/privatesquare/bkst-oauth-api/interfaces/db/cassandra"
)

func NewAccessTokenService(r cassandra.AccessTokenStore) AccessTokenService {
	return &accessTokenService{
		AccessTokenStore: r,
	}
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type accessTokenService struct {
	AccessTokenStore cassandra.AccessTokenStore
}

func (ats *accessTokenService) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return ats.AccessTokenStore.GetById(id)
}
