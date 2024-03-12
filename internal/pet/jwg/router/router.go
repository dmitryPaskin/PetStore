package router

import (
	"github.com/go-chi/chi"
	"studentgit.kata.academy/xp/PetStore/internal/pet/controller"
	"studentgit.kata.academy/xp/PetStore/internal/pet/service"
)

func NewRouter(r chi.Router, ps service.PetService) {
	p := controller.New(ps)
	r.Route("/pet", func(r chi.Router) {
		r.Get("/{petId}", p.Get)
		r.Post("/", p.Create)
		r.Put("/", p.Update)
		r.Delete("/{petId}", p.Delete)
		r.Get("/findByStatus", p.FindByStatus)
	})
}
