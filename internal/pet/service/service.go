package service

import (
	"context"
	"studentgit.kata.academy/xp/PetStore/internal/pet/entities"
	"studentgit.kata.academy/xp/PetStore/internal/pet/repository"
)

type PetService interface {
	GetByStatusS(ctx context.Context, status string) ([]*entities.Pet, error)
	GetS(ctx context.Context, id int) (*entities.Pet, error)
	UpdateS(ctx context.Context, pet *entities.Pet) error
	DeleteS(ctx context.Context, id int) error
	CreateS(ctx context.Context, pet *entities.Pet) error
}

type petService struct {
	petRepo repository.PetRepository
}

func New(pr repository.PetRepository) PetService {
	return &petService{
		pr,
	}
}

func (p *petService) GetByStatusS(ctx context.Context, status string) ([]*entities.Pet, error) {
	return p.petRepo.GetByStatus(ctx, status)
}

func (p *petService) GetS(ctx context.Context, id int) (*entities.Pet, error) {
	return p.petRepo.Get(ctx, id)
}

func (p *petService) UpdateS(ctx context.Context, pet *entities.Pet) error {
	return p.petRepo.Update(ctx, pet)
}

func (p *petService) DeleteS(ctx context.Context, id int) error {
	return p.petRepo.Delete(ctx, id)
}

func (p *petService) CreateS(ctx context.Context, pet *entities.Pet) error {
	return p.CreateS(ctx, pet)
}
