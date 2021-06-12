package users_api

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/privatesquare/bkst-oauth-api/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var (
	usersApiBaseUrl = "https://bkst.example.com"

	login = domain.Login{
		Username: "hpotter@hogwarts.com",
		Password: "password",
	}

	userResp = domain.User{
		Id:        1,
		FirstName: "Harry",
		Lastname:  "Potter",
		Email:     "hpotter@hogwarts.com",
		Status:    "active",
	}

	userNotFoundRestErr       = errors.NotFoundError("User with email notfound@hogwarts.com was not found")
	invalidCredentialsRestErr = errors.BadRequestError("username or password is incorrect")
)

func init() {
	userApiCfg := UsersApiCfg{Url: usersApiBaseUrl}
	userApiCfg.SetConfig()
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNewUsersStore(t *testing.T) {
	us := NewUsersStore(resty.New())
	assert.NotNil(t, us)
}

func TestUsersStore_Login(t *testing.T) {
	client := resty.New().SetHostURL(baseUrl)
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder, err := httpmock.NewJsonResponder(http.StatusOK, &userResp)
	assert.NoError(t, err)
	httpmock.RegisterResponder("POST", baseUrl+usersLoginApiPath, responder)

	us := NewUsersStore(client)
	user, restErr := us.Login(login)

	assert.Nil(t, restErr)
	assert.Equal(t, userResp.Id, user.Id)
	assert.Equal(t, userResp.FirstName, user.FirstName)
	assert.Equal(t, userResp.Lastname, user.Lastname)
	assert.Equal(t, userResp.Email, user.Email)
	assert.Equal(t, userResp.Status, user.Status)
}

func TestUsersStore_Login_RequestError(t *testing.T) {
	client := resty.New().SetHostURL(baseUrl)
	us := NewUsersStore(client)
	user, restErr := us.Login(login)

	assert.Nil(t, user)
	assert.Equal(t, http.StatusInternalServerError, restErr.Status)
}

func TestUsersStore_Login_UserNotFound(t *testing.T) {
	client := resty.New().SetHostURL(baseUrl)
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder, err := httpmock.NewJsonResponder(http.StatusNotFound, &userNotFoundRestErr)
	assert.NoError(t, err)
	httpmock.RegisterResponder("POST", baseUrl+usersLoginApiPath, responder)

	us := NewUsersStore(client)
	user, restErr := us.Login(login)

	assert.Nil(t, user)
	assert.Equal(t, userNotFoundRestErr.Message, restErr.Message)
	assert.Equal(t, userNotFoundRestErr.Status, restErr.Status)
	assert.Equal(t, userNotFoundRestErr.Error, restErr.Error)
}

func TestUsersStore_Login_InvalidCredentials(t *testing.T) {
	client := resty.New().SetHostURL(baseUrl)
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder, err := httpmock.NewJsonResponder(http.StatusNotFound, &invalidCredentialsRestErr)
	assert.NoError(t, err)
	httpmock.RegisterResponder("POST", baseUrl+usersLoginApiPath, responder)

	us := NewUsersStore(client)
	user, restErr := us.Login(login)

	assert.Nil(t, user)
	assert.Equal(t, invalidCredentialsRestErr.Message, restErr.Message)
	assert.Equal(t, invalidCredentialsRestErr.Status, restErr.Status)
	assert.Equal(t, invalidCredentialsRestErr.Error, restErr.Error)
}
