package router

import (
	"github.com/go-chi/chi"
	"studentgit.kata.academy/xp/PetStore/internal/store/controller"
	"studentgit.kata.academy/xp/PetStore/internal/store/service"
)

func NewRouter(r chi.Router, ss service.StoreService) {
	s := controller.New(ss)

	r.Route("/store/order", func(r chi.Router) {
		r.Get("/{orderId}", s.Get)
		r.Delete("/{orderId}", s.Delete)
		r.Post("/", s.Create)
	})
}
