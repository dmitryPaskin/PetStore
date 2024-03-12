package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"studentgit.kata.academy/xp/PetStore/internal/user/entities"
	"studentgit.kata.academy/xp/PetStore/internal/user/service"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	CreateWithList(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service.UserService
}

func New(us service.UserService) UserController {
	return &userController{
		us,
	}
}

// @Summary Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body entities.User true "User to add"
// @Success 200 {string} string "User created"
// Router /user [post]
func (u *userController) Create(w http.ResponseWriter, r *http.Request) {
	var userInput entities.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.UserService.Create(r.Context(), &userInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
	if err != nil {
		return
	}
}

// @Summary Get user by username
// @Tags user
// @Produce json
// @Param username path string true "Username of user to return"
// @Success 200 {object} entities.User "Find user by username"
// @Router /user/{username} [get]
func (u *userController) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, fmt.Errorf("param user is not set").Error(), http.StatusBadRequest)
		return
	}

	user, err := u.UserService.Get(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Update a user with form data
// @Tags user
// @Accept json
// @Produce json
// @Param user body entities.User true "User object that needs to update"
// @Success 200 {string} string "User updated"
// @Router /user/{username} [put]
func (u *userController) Update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, fmt.Errorf("param user is not set").Error(), http.StatusBadRequest)
		return
	}

	var userInput entities.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.UserService.Update(r.Context(), username, &userInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "User update"})
	if err != nil {
		return
	}
}

// @Summary	Delete a user by username
// @Tags user
// @Produce	json
// @Param username path string true	"Username of user to delete"
// @Success	 200 {string} string "User deleted"
// @Router /user/{username} [delete]
func (u *userController) Delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, fmt.Errorf("param user is not set").Error(), http.StatusBadRequest)
		return
	}

	if err := u.UserService.Delete(r.Context(), username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
	if err != nil {
		return
	}
}

// @Summary	Create a list of new users
// @Tags user
// @Accept	json
// @Produce	json
// @Param users	body []entities.User true "Users to add to the store"
// @Success	200	 {string} string	"Users created"
// @Router	/user/createWithList [post]
func (u *userController) CreateWithList(w http.ResponseWriter, r *http.Request) {
	var userInput []*entities.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.UserService.CreateList(r.Context(), userInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Users created"})
	if err != nil {
		return
	}
}

// @Summary	Login a user
// @Tags user
// @Accept json
// @Produce	json
// @Param credentials body	entities.LoginRequest true "User credentials"
// @Success	200	{string} string	"User login"
// @Router /user/login	[get]
func (u *userController) Login(w http.ResponseWriter, r *http.Request) {
	var loginInput entities.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := u.UserService.Login(r.Context(), loginInput.Username, loginInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		return
	}
}

// @Summary	Logout a user
// @Tags user
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Success	200	{string} string	"User logout"
// @Router	/user/logout [get]
func (u *userController) Logout(w http.ResponseWriter, r *http.Request) {
	token := jwtauth.TokenFromHeader(r)
	if token == "" {
		http.Error(w, fmt.Errorf("token is not set").Error(), http.StatusBadRequest)
		return
	}

	if err := u.UserService.Logout(r.Context(), token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "User logout"})
	if err != nil {
		return
	}
}
