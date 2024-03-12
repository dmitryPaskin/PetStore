package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"studentgit.kata.academy/xp/PetStore/internal/pet/entities"
	"studentgit.kata.academy/xp/PetStore/internal/pet/service"
)

type PetController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindByStatus(w http.ResponseWriter, r *http.Request)
}

type petController struct {
	petService service.PetService
}

func New(ps service.PetService) PetController {
	return &petController{
		petService: ps,
	}
}

// @Summary	Add a new pet to the store
// @Tags pet
// @Accept json
// @Produce	json
// @Security  ApiKeyAuth
// @Param pet body entities.Pet	true "Pet object that needs to be added to the store"
// @Success 200	{object} entities.Pet "Pet object that was added"
// @Router /pet [post]
func (p *petController) Create(w http.ResponseWriter, r *http.Request) {
	var petInput entities.Pet

	if err := json.NewDecoder(r.Body).Decode(&petInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.petService.CreateS(r.Context(), &petInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "pet created"})
	if err != nil {
		return
	}
}

// @Summary	Get a pet by ID
// @Tags pet
// @Produce	json
// @Security ApiKeyAuth
//
//	@Param petId path int true "ID of pet to return"
//
// @Success	200	{object} entities.Pet "Find pet by ID"
// @Router /pet/{petId} [get]
func (p *petController) Get(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	if petId == "" {
		http.Error(w, fmt.Errorf("param petId is required").Error(), http.StatusBadRequest)
		return
	}

	petIdInt, err := strconv.Atoi(petId)
	if err != nil {
		http.Error(w, fmt.Errorf("param petId must by integer").Error(), http.StatusBadRequest)
		return
	}

	pet, err := p.petService.GetS(r.Context(), petIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary	Update a pet in the store with form data
// @Tags pet
// @Accept json
// @Produce	json
// @Security ApiKeyAuth
// @Param pet body entities.Pet true "Pet object that needs to update"
// @Success 200 {string} string "Pet updated"
// @Router /pet [put]
func (p *petController) Update(w http.ResponseWriter, r *http.Request) {
	var petInput entities.Pet

	if err := json.NewDecoder(r.Body).Decode(&petInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.petService.UpdateS(r.Context(), &petInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "pet updated"})
	if err != nil {
		return
	}
}

// @Summary	Delete a pet by ID
// @Tags pet
// @Produce	json
// @Security ApiKeyAuth
// @Param petId	path int true "ID of pet to delete"
// @Success	200	{string} string	"Pet deleted"
// @Router /pet/{petId} [delete]
func (p *petController) Delete(w http.ResponseWriter, r *http.Request) {
	petId := chi.URLParam(r, "petId")
	if petId == "" {
		http.Error(w, fmt.Errorf("param petId is required").Error(), http.StatusBadRequest)
		return
	}

	petIdInt, err := strconv.Atoi(petId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = p.petService.DeleteS(r.Context(), petIdInt); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "pet deleted"})
	if err != nil {
		return
	}
}

// @Summary	Delete a pet by ID
// @Tags pet
// @Produce	json
// @Security ApiKeyAuth
// @Param status query string true "Status values that need to be considered for filter"
// @Success	200	{object} []entities.Pet "Pets found by status"
// @Router /pet/findByStatus [get]
func (p *petController) FindByStatus(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("status") {
		http.Error(w, fmt.Errorf("param status is required").Error(), http.StatusBadRequest)
		return
	}

	status := r.URL.Query().Get("status")
	pets, err := p.petService.GetByStatusS(r.Context(), status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(pets); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
