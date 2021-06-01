package db

import (
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-oauth-api/domain/access_token"
)

func New() Repository {
	return &repository{}
}

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type repository struct {}

func (r *repository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.InternalServerError("db repository is not implemented yet!")
}
