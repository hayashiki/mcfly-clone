package api_test

import (
	"testing"

	"github.com/hayashiki/mcfly-clone/api"
	"github.com/hayashiki/mcfly-clone/client"
)

type EndpointTest struct {
	JSON         string
	Desc         string
	Should       string
	ExpectStatus int
	ExpectBody   interface{}
	AccessToken  string
}

func RunEndpointTests(t *testing.T, httpMethod string, endpointPath string, tests []*EndpointTest) {
	for _, et := range tests {

		var jsonData *string
		if et.JSON == "" {
			jsonData = nil
		} else {
			jsonData = &et.JSON
		}

		reqOpts := &client.RequestOptions{}
		if et.AccessToken != "" {
			reqOpts.AuthHeader = &et.AccessToken
		}

		res, err := apiClient.DoReq(httpMethod, endpointPath, jsonData, reqOpts)
		if err != nil {
			t.Error(err)
		}
		defer res.Body.Close()
	}
}

func MissingAuthorizationHeaderEndpointTest(json string) *EndpointTest {
	return &EndpointTest{
		json,
		"nothing in the Authorization header",
		"an authorization header required error",
		400,
		api.NewAuthorizationHeaderRequiredErr(),
		"",
	}
}

func InvalidAuthorizationTokenErrorTest(json string) *EndpointTest {
	return &EndpointTest{
		json,
		"an invalid access token",
		"an invalid access token error",
		400,
		api.NewInvalidAuthTokenError("mock_invalid_access_token"),
		"mock_invalid_access_token",
	}
}
