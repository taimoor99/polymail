package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"polymail/app/repository"
)

func NewRouter(ctrl repository.DraftMail) *chi.Mux {
	r := chi.NewRouter()
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Group(func(r chi.Router) {
		r.Post("/addmaildraft", ctrl.CreateDraftMailHandler)
		r.Get("/getmaildraft/{id}", ctrl.GetDraftMailHandler)
		r.Delete("/deletemaildraft/{id}", ctrl.DeleteDraftMailHandler)
		r.Put("/updatemaildraft/{id}", ctrl.UpdateDraftMailHandler)
		r.Put("/senddraftemail/{id}", ctrl.SendDraftMailHandler)
	})

	return r
}
