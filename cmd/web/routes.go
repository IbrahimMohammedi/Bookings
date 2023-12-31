package main

import (
	"net/http"

	"github.com/IbrahimMohammedi/Bookings/internal/config"
	"github.com/IbrahimMohammedi/Bookings/internal/handlers"
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
	mux.Post("/booking", handlers.Repo.PostBook)
	mux.Post("/bookingJSON", handlers.Repo.BookJSON)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservations", handlers.Repo.Reservation)
	mux.Post("/make-reservations", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Route("/admin", func(r chi.Router) {
		mux.Use(Authed)

		mux.Get("/dashboard", handlers.Repo.AdminDashBoard)
	})

	return mux
}
