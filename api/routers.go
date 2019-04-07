package api

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func AllRoutes(handlers *Handlers) Routes {
	return Routes{
		Route{
			"GetProviderProjects",
			"GET",
			"/api/0/projects",
			handlers.MakeHandlerFunc(HandlerOptions{
				AuthRequired: true,
				After:        handlers.GetProjects,
			}),
		},
	}
}
