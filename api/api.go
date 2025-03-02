package main

import (
	"net/http"
	repository "zviziviso-app/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Application struct {
	config  applicationConfig
	service repository.Querier
}

type applicationConfig struct {
	address string
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(render.SetContentType(render.ContentTypeHTML))

	r.Route("/v1", func(r chi.Router) {

		r.Route("/countries", func(r chi.Router) {
			r.Get("/", app.ListCountries)
			r.Route("/{countryId}", func(r chi.Router) {
				r.Get("/", app.GetCountryById)
				r.Get("/provinces", app.GetProvincesByCountryId)
			})
		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/", app.CreateNewUser)
			r.Get("/{userId}", app.GetUserSummaryDetails)
		})

		r.Route("/death_notice", func(r chi.Router) {
			r.Get("/", app.NewDeathNotice)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}

func (app *Application) Run(mux http.Handler) error {

	server := http.Server{
		Addr:    app.config.address,
		Handler: mux,
	}

	return server.ListenAndServe()
}
