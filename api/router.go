package api

import (
	"github.com/hayashiki/mcfly-clone/api"
	"github.com/gorilla/mux"
	// "github.com/hayashiki/mcfly-clone/api/apidata"
)

type RequestLogger interface {
	Handler(http.Handler, string) http.Handler
}

func NewRouter(
	cfg *config.Config
)
	*mux.Router {
		db, err := models.NewDB(cfg.DatabaseUrl, cfg.DatabaseName, vfg.DatabaesUseSSL)
		if err != nil {
			log.Panic(err)
		}

		handlers := &Handlers{
			DB: db,
		}

		log.Printf("connected to postgres")

		router := mux.NewRouter()

		routes := AllRoutes(handlers)

		for _, route := routes {
			var handler http.Handler

			handler = logger.Hanlder{route.HanlderFunc, route.Name}

			router.
				Methods(route.Methods).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}

		return router

	}