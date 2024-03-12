package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"studentgit.kata.academy/xp/PetStore/internal/jwg/dataBase"
	"studentgit.kata.academy/xp/PetStore/internal/jwg/superRouter"
	"time"
)

// @titile PetStore
// @version 1.0
// @description Implementation of PetStore API
// @securitydefinitions.apikey ApiKeyAuth
// @in Login
// @name Authorization
// @host localhost:8080
// @BasePath /
func RunApp() {
	db, err := dataBase.New()
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))

	r := chi.NewRouter()

	router := superRouter.New(r, db.DB, tokenAuth)
	router.StartRouter()
}
