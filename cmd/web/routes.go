package main

import (
	"net/http"

	"github.com/IbrahimMohammedi/Bookings/pkg/config"
	"github.com/IbrahimMohammedi/Bookings/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/booking", handlers.Repo.Book)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservations", handlers.Repo.Reservations)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	return mux
}
