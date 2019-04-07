package api_test

import (
	"testing"

	"github.com/hayashiki/mcfly-clone/api/apidata"
)

func TestGetProject(t *testing.T) {
	tests := []*EndpointTest{
		MissingAuthorizationHeaderEndpointTest(""),
		InvalidAuthorizationTokenErrorTest(""),
		{
			"",
			"a request to get projects added to repofly",
			"a list of mcfly projects",
			200,
			[]apidata.ProjectResp{
				{"mattmocks/project1", "https://jabroni.com/mattmocks/project-1", "jabroni.com"},
				{"mattmocks/project2", "https://jabroni.com/mattmocks/project-2", "jabroni.com"},
				{"mattmocks/project3", "https://jabroni.com/mattmocks/project-3", "jabroni.com"},
			},
			"mock_seede_access_token_123",
		},
	}
	RunEndpointTests(t, "GET", "projects", tests)
}
