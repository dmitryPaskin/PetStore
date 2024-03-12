package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"studentgit.kata.academy/xp/PetStore/internal/store/entities"
	"studentgit.kata.academy/xp/PetStore/internal/store/service"
)

type StoreController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type storeController struct {
	storeServ service.StoreService
}

func New(ss service.StoreService) StoreController {
	return &storeController{
		storeServ: ss,
	}
}

// @Summary	Add a new order to the store
// @Tags 	store
// @Accept	json
// @Produce	json
// @Security ApiKeyAuth
// @Param	order body	entities.Store true	"Order object that needs to be added to the store"
// @Success	200	{object} entities.Store "Order object that was added"
// @Router /store/order [post]
func (s *storeController) Create(w http.ResponseWriter, r *http.Request) {
	var storeInput entities.Store
	if err := json.NewDecoder(r.Body).Decode(&storeInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.storeServ.CreateS(r.Context(), &storeInput); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Store created"})
	if err != nil {
		return
	}
}

// @Summary	Order an order by ID
// @Tags store
// @Produce	json
// @Security ApiKeyAuth
// @Param orderId path int	true "ID of order to return"
// @Success	200	{object} entities.Store "Find order by ID"
// @Router /store/order/{orderId} [get]
func (s *storeController) Get(w http.ResponseWriter, r *http.Request) {
	storeIdStr := chi.URLParam(r, "orderId")
	if storeIdStr == "" {
		http.Error(w, fmt.Errorf("orderId is not set").Error(), http.StatusBadRequest)
		return
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store, err := s.storeServ.GetS(r.Context(), storeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(store); err != nil {
		return
	}
}

// @Summary	Delete order by ID
// @Tags store
// @Produce	json
// @Security ApiKeyAuth
// @Param orderId path int	true "ID of order to delete"
// @Success	200	{string} string	"Order deleted"
// @Router /store/order/{orderId} [delete]
func (s *storeController) Delete(w http.ResponseWriter, r *http.Request) {
	storeIdStr := chi.URLParam(r, "orderId")
	if storeIdStr == "" {
		http.Error(w, fmt.Errorf("orderId is not set").Error(), http.StatusBadRequest)
		return
	}

	storeId, err := strconv.Atoi(storeIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.storeServ.DeleteS(r.Context(), storeId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(map[string]string{"message": "Store deleted"}); err != nil {
		return
	}
}
