package service

import (
	"context"
	"studentgit.kata.academy/xp/PetStore/internal/store/entities"
	"studentgit.kata.academy/xp/PetStore/internal/store/repository"
)

type StoreService interface {
	CreateS(ctx context.Context, store *entities.Store) error
	GetS(ctx context.Context, id int) (*entities.Store, error)
	DeleteS(ctx context.Context, id int) error
}

type storeService struct {
	storeRepo repository.StoreRepository
}

func New(sr repository.StoreRepository) StoreService {
	return &storeService{
		storeRepo: sr,
	}
}

func (s *storeService) CreateS(ctx context.Context, store *entities.Store) error {
	return s.storeRepo.Create(ctx, store)
}

func (s *storeService) GetS(ctx context.Context, id int) (*entities.Store, error) {
	return s.storeRepo.Get(ctx, id)
}

func (s *storeService) DeleteS(ctx context.Context, id int) error {
	return s.storeRepo.Delete(ctx, id)
}
