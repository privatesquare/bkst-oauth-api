package services

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-oauth-api/domain"
	"github.com/privatesquare/bkst-oauth-api/interfaces/db/cassandra"
)

// TODO: Add validation

func NewAccessTokenService(r cassandra.AccessTokenStore) AccessTokenService {
	return &accessTokenService{
		AccessTokenStore: r,
	}
}

type AccessTokenService interface {
	GetById(string) (*domain.AccessToken, *errors.RestErr)
	Create(domain.AccessToken) *errors.RestErr
	Update(domain.AccessToken) *errors.RestErr
}

type accessTokenService struct {
	AccessTokenStore cassandra.AccessTokenStore
}

func (ats *accessTokenService) GetById(id string) (*domain.AccessToken, *errors.RestErr) {
	return ats.AccessTokenStore.GetById(id)
}

func (ats *accessTokenService) Create(at domain.AccessToken) *errors.RestErr {
	at.SetExpiration()
	if err := at.Validate(); err != nil {
		return err
	}
	return ats.AccessTokenStore.Create(at)
}

func (ats *accessTokenService) Update(at domain.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return ats.AccessTokenStore.Update(at)
}
