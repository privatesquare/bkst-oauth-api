package domain

import (
	"github.com/privatesquare/bkst-go-utils/utils/dateutils"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-go-utils/utils/structutils"
	"strings"
	"time"
)

const (
	defaultExpirationTimeInSeconds = 90
)

type AccessToken struct {
	AccessToken string    `json:"access_token"`
	ClientId    int64     `json:"client_id"`
	UserId      int64     `json:"user_id"`
	Expires     time.Time `json:"expires"`
}

func (at *AccessToken) SetExpiration() {
	at.Expires = dateutils.GetDateTimeNow().Add(defaultExpirationTimeInSeconds * time.Second)
}

func (at *AccessToken) Validate() *errors.RestErr {
	return at.validateFields()
}

func (at *AccessToken) validateFields() *errors.RestErr {
	var missingParams []string

	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		missingParams = append(missingParams, structutils.GetFieldTagValue(at, &at.AccessToken))
	}
	if at.ClientId <= 0 {
		missingParams = append(missingParams, structutils.GetFieldTagValue(at, &at.ClientId))
	}
	if at.UserId <= 0 {
		missingParams = append(missingParams, structutils.GetFieldTagValue(at, &at.UserId))
	}
	if at.Expires.IsZero() {
		missingParams = append(missingParams, structutils.GetFieldTagValue(at, &at.Expires))
	}

	if len(missingParams) > 0 {
		err := errors.MissingMandatoryParamError(missingParams)
		return errors.BadRequestError(err.Error())
	}
	return nil
}

func (at *AccessToken) IsExpired() bool {
	return dateutils.GetDateTimeNow().After(at.Expires)
}
