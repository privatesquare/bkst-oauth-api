package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"github.com/privatesquare/bkst-oauth-api/domain"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateAccessToken = "UPDATE access_tokens SET expires=? WHERE access_tokens=?;"

	accessTokenNotFoundMsg = "Access token with id %s was not found"
	dbCreateSessionErrMsg  = "Unable to create db session"
	dbQueryErrMsg          = "Unable to execute database query"
)

func NewAccessTokenStore() AccessTokenStore {
	return &accessTokenStore{}
}

type AccessTokenStore interface {
	GetById(string) (*domain.AccessToken, *errors.RestErr)
	Create(domain.AccessToken) *errors.RestErr
	Update(domain.AccessToken) *errors.RestErr
}

type accessTokenStore struct{}

func (ats *accessTokenStore) GetById(id string) (*domain.AccessToken, *errors.RestErr) {
	at := new(domain.AccessToken)
	if err := session.Query(queryGetAccessToken, id).Scan(&at.AccessToken,
		&at.UserId, &at.ClientId, &at.Expires); err != nil {
		if err == gocql.ErrNotFound {
			msg := fmt.Sprintf(accessTokenNotFoundMsg, id)
			logger.Info(msg)
			return nil, errors.NotFoundError(msg)
		}
		logger.Error(dbQueryErrMsg, err)
		return nil, errors.InternalServerError()
	}
	return at, nil
}

func (ats *accessTokenStore) Create(at domain.AccessToken) *errors.RestErr {
	if err := session.Query(queryCreateAccessToken, &at.AccessToken, &at.UserId,
		&at.ClientId, &at.Expires).Exec(); err != nil {
		logger.Error(dbQueryErrMsg, err)
		return errors.InternalServerError()
	}
	return nil
}

func (ats *accessTokenStore) Update(at domain.AccessToken) *errors.RestErr {
	if err := session.Query(queryUpdateAccessToken, &at.Expires, &at.AccessToken).Exec(); err != nil {
		logger.Error(dbQueryErrMsg, err)
		return errors.InternalServerError()
	}
	return nil
}
