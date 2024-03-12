package service

import (
	"context"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"studentgit.kata.academy/xp/PetStore/internal/user/entities"
	"studentgit.kata.academy/xp/PetStore/internal/user/repository"
)

type UserService interface {
	Create(ctx context.Context, user *entities.User) error
	Get(ctx context.Context, username string) (*entities.User, error)
	Update(ctx context.Context, username string, user *entities.User) error
	Delete(ctx context.Context, username string) error
	CreateList(ctx context.Context, users []*entities.User) error
	Login(ctx context.Context, username string, password string) (string, error)
	Logout(ctx context.Context, jwtToken string) error
}

type userService struct {
	userRepo repository.UserRepository
	jwtAuth  *jwtauth.JWTAuth
}

func New(ur repository.UserRepository, jwt *jwtauth.JWTAuth) UserService {
	return &userService{
		userRepo: ur,
		jwtAuth:  jwt,
	}
}

func (u *userService) Create(ctx context.Context, user *entities.User) error {
	user.Password = hashPassword(user.Password)
	return u.userRepo.Create(ctx, user)
}

func (u *userService) Get(ctx context.Context, username string) (*entities.User, error) {
	return u.userRepo.GetByUsername(ctx, username)
}

func (u *userService) Update(ctx context.Context, username string, user *entities.User) error {
	user.Password = hashPassword(user.Password)
	return u.userRepo.Update(ctx, username, user)
}

func (u *userService) Delete(ctx context.Context, username string) error {
	userId, err := u.userRepo.GetIdByUsername(ctx, username)
	if err != nil {
		return err
	}

	if err := u.userRepo.Logout(ctx, userId); err != nil {
		return err
	}

	return u.userRepo.Delete(ctx, username)
}

func (u *userService) CreateList(ctx context.Context, users []*entities.User) error {
	for _, user := range users {
		err := u.userRepo.Create(ctx, user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *userService) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := u.Get(ctx, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	sessionId, err := u.userRepo.Login(ctx, user.Id)
	if err != nil {
		return "", err
	}

	_, jwtToken, _ := u.jwtAuth.Encode(map[string]interface{}{"session_id": sessionId})
	return jwtToken, nil
}

func (u *userService) Logout(ctx context.Context, jwtToken string) error {
	decodeJwtToken, err := u.jwtAuth.Decode(jwtToken)
	if err != nil {
		return err
	}

	sessionId, ok := decodeJwtToken.Get("session_id")
	if !ok {
		return fmt.Errorf("session_id not found")
	}

	return u.userRepo.Logout(ctx, int(sessionId.(float64)))
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
