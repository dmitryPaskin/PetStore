package router

import (
	"github.com/go-chi/chi"
	"studentgit.kata.academy/xp/PetStore/internal/user/controller"
	"studentgit.kata.academy/xp/PetStore/internal/user/service"
)

func NewRouter(r chi.Router, us service.UserService) {
	u := controller.New(us)
	r.Route("/user", func(r chi.Router) {
		r.Get("/login", u.Login)
		r.Get("/login", u.Logout)
		r.Get("/{username}", u.Get)

		r.Post("/createWithList", u.CreateWithList)
		r.Post("/", u.Create)

		r.Put("/{username}", u.Update)

		r.Delete("/{username}", u.Delete)
	})
}
