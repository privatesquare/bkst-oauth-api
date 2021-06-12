package users_api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
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

	user = domain.User{
		Id:        1,
		FirstName: "Harry",
		Lastname:  "Potter",
		Email:     "hpotter@hogwarts.com",
		Status:    "active",
	}
)

func init() {
	userApiCfg := UsersApiCfg{Url: usersApiBaseUrl}
	userApiCfg.SetConfig()
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestUsersStore_Login_ResponseErr(t *testing.T) {
	client := resty.New().SetHostURL(baseUrl)
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder, err := httpmock.NewJsonResponder(http.StatusOK, &user)
	assert.NoError(t, err)
	httpmock.RegisterResponder("POST", baseUrl+usersLoginApiPath, responder)

	us := NewUsersStore(client)
	user, restErr := us.Login(login)

	fmt.Println(user)
	fmt.Println(restErr)
}
