package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kyong0612/otp-getter/handler"
)

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handler.GetReadOtpPage)
	r.Post("/otp", handler.ReadOtpHandler)

	return r
}
