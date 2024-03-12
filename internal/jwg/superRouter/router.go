package superRouter

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"studentgit.kata.academy/xp/PetStore/internal/middleware"
	"syscall"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	_userRouter "studentgit.kata.academy/xp/PetStore/internal/user/jwg/router"
	_userRepo "studentgit.kata.academy/xp/PetStore/internal/user/repository"
	_userServ "studentgit.kata.academy/xp/PetStore/internal/user/service"

	_petRouter "studentgit.kata.academy/xp/PetStore/internal/pet/jwg/router"
	_petRepo "studentgit.kata.academy/xp/PetStore/internal/pet/repository"
	_petServ "studentgit.kata.academy/xp/PetStore/internal/pet/service"

	_storeRouter "studentgit.kata.academy/xp/PetStore/internal/store/jwg/router"
	_storeRepo "studentgit.kata.academy/xp/PetStore/internal/store/repository"
	_storeServ "studentgit.kata.academy/xp/PetStore/internal/store/service"

	_ "studentgit.kata.academy/xp/PetStore/internal/docs"
)

type SuperRouter struct {
	chi *chi.Mux
	db  *sql.DB
	jwt *jwtauth.JWTAuth
}

func New(chi *chi.Mux, db *sql.DB, jwt *jwtauth.JWTAuth) SuperRouter {
	return SuperRouter{
		chi: chi,
		db:  db,
		jwt: jwt,
	}
}

func (sr *SuperRouter) StartRouter() {

	sr.chi.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	sr.chi.Group(func(r chi.Router) {
		userRepo := _userRepo.New(sr.db)
		userServ := _userServ.New(userRepo, sr.jwt)

		_userRouter.NewRouter(r, userServ)
	})

	sr.chi.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(sr.jwt))
		r.Use(middleware.AuthMiddleware)

		petRepo := _petRepo.New(sr.db)
		petServ := _petServ.New(petRepo)

		_petRouter.NewRouter(r, petServ)
	})

	sr.chi.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(sr.jwt))
		r.Use(middleware.AuthMiddleware)

		storeRepo := _storeRepo.New(sr.db)
		storeServ := _storeServ.New(storeRepo)

		_storeRouter.NewRouter(r, storeServ)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      sr.chi,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigChan
	stopCTX, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(stopCTX); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
