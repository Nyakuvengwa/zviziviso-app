package main

import (
	"context"
	"log"
	"net/http"
	application_constants "zviziviso-app/internal/constants/application_constants"
	repository "zviziviso-app/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
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
			r.Use(app.UserRequiredMiddleware)
			r.Post("/", app.CreateNewDeathNotice)
			r.Get("/{deathNoticeId}", app.GetDeathNoticeById)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})


	return r
}

func (app *Application) UserRequiredMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId := r.Header.Get("X-UserId")
		if userId == "" {
			NewProblemDetailsErrorResponse(w, http.StatusBadRequest, "Invalid user id provided", "Invalid header provided.")
			return
		}

		// TODO: Add a layer of caching this user to avoid overloading the user
		// Research if Go has in memory cache
		user, err := app.service.GetUserSummaryDetails(ctx, uuid.MustParse(userId))

		if err != nil {
			NewProblemDetailsErrorResponse(w, http.StatusInternalServerError, "Unhandled server error", err.Error())
			return
		}

		if user.UserID == uuid.Nil {
			NewProblemDetailsErrorResponse(w, http.StatusNotFound, "User not found", "User not found.")
			return
		}

		log.Printf("User Found: %s,", user.UserID)
		ctx = context.WithValue(ctx, application_constants.User, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func (app *Application) Run(mux http.Handler) error {

	server := http.Server{
		Addr:    app.config.address,
		Handler: mux,
	}

	return server.ListenAndServe()
}
