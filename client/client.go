package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client interface {
}

type ClientResponse struct {
	Data       interface{}
	StatusCode int
}

type McflyClient struct {
	ServerURL   string
	AccessToken string
}

type RequestOptions struct {
	UseBasicAuth bool
	AuthHeader   *string
}

func (self *McflyClient) DoReq(method string, endpoint string, JSON *string, opts *RequestOptions) (*http.Response, error) {
	if opts == nil {
		opts = NewRequestOptions()
	}

	url := self.EndpointURL(endpoint)

	var reader io.Reader
	if JSON != nil {
		reader = strings.NewReader(*JSON)
	}

	req, err := http.NewRequest(method, url, reader)

	if err != nil {
		return nil, err
	}

	if self.AccessToken != "" && opts.AuthHeader == nil {
		opts.AuthHeader = &self.AccessToken
	}

	if opts.AuthHeader != nil {
		req.Header.Add("Authorization", *opts.AuthHeader)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (self *McflyClient) EndpointURL(endpointName string) string {
	return fmt.Sprintf("%s/api/0/%s", self.ServerURL, endpointName)
}

func NewRequestOptions() *RequestOptions {
	return &RequestOptions{true, nil}
}
