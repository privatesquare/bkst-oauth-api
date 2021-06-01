package services

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-oauth-api/domain/access_token"
	"github.com/privatesquare/bkst-oauth-api/repository/db"
)

func New(r db.Repository) AccessTokenService {
	return &accessTokenService{
		Repository: r,
	}
}

type AccessTokenService interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type accessTokenService struct {
	Repository db.Repository
}

func (ats *accessTokenService) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return ats.Repository.GetById(id)
}
