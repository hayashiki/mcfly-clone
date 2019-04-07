package mockclient

import (
	"errors"
	"net/http"

	"github.com/hayashiki/mcfly-clone/api/apidata"
	"github.com/hayashiki/mcfly-clone/client"
	"github.com/stretchr/testify/mock"
)

var ErrMock = errors.New("mock error")

var ClientResponseError = &client.ClientResponse{
	StatusCode: 400,
	Data:       &apidata.ApiError{Error: "mock api error"},
}

var ClientResponseSuccess = &client.ClientResponse{
	StatusCode: 200,
	Data:       &apidata.ApiResponse{"success!"},
}

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Login(request *apidata.LoginReq) (*client.ClientResponse, *http.Response, error) {
	args := c.Called(request)
	return args.Get(0).(*client.ClientResponse), args.Get(1).(*http.Response), args.Error(2)
}
