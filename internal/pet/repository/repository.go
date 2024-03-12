package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"studentgit.kata.academy/xp/PetStore/internal/pet/entities"
)

type PetRepository interface {
	GetByStatus(ctx context.Context, status string) ([]*entities.Pet, error)
	Get(ctx context.Context, id int) (*entities.Pet, error)
	Update(ctx context.Context, pet *entities.Pet) error
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, pet *entities.Pet) error
}

type petRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(db *sql.DB) PetRepository {
	return &petRepository{
		db:         db,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (p *petRepository) GetByStatus(ctx context.Context, status string) ([]*entities.Pet, error) {
	query := p.sqlBuilder.Select("id", "category", "name", "tags", "status").
		From("pets").Where(sq.Eq{"status": status})

	rows, err := query.RunWith(p.db).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pets := make([]*entities.Pet, 0)

	for rows.Next() {
		var pet entities.Pet
		if err := rows.Scan(&pet.Id, &pet.Category, &pet.Name, &pet.Tags, &pet.Status); err != nil {
			return nil, err
		}

		pets = append(pets, &pet)
	}
	return pets, nil
}

func (p *petRepository) Get(ctx context.Context, id int) (*entities.Pet, error) {
	query := p.sqlBuilder.Select("id", "category", "name", "tags", "status").
		From("pets").Where(sq.Eq{"id": id})

	row := query.RunWith(p.db).QueryRowContext(ctx)
	var pet entities.Pet
	if err := row.Scan(&pet.Id, &pet.Category, &pet.Name, &pet.Tags, &pet.Status); err != nil {
		return nil, err
	}

	return &pet, nil
}

func (p *petRepository) Update(ctx context.Context, pet *entities.Pet) error {
	query := p.sqlBuilder.Update("pets").
		Set("category", pet.Category).
		Set("name", pet.Name).
		Set("tags", pet.Tags).
		Set("status", pet.Status).
		Where(sq.Eq{"id": pet.Id})

	if _, err := query.RunWith(p.db).ExecContext(ctx); err != nil {
		return err
	}
	return nil
}

func (p *petRepository) Delete(ctx context.Context, id int) error {
	query := p.sqlBuilder.Delete("pets").Where(sq.Eq{"id": id})
	if _, err := query.RunWith(p.db).ExecContext(ctx); err != nil {
		return err
	}

	return nil
}

func (p *petRepository) Create(ctx context.Context, pet *entities.Pet) error {
	query := p.sqlBuilder.Insert("pets").Columns("category", "name", "tags", "status").
		Values(pet.Category, pet.Name, pet.Tags, pet.Status).
		Suffix("RETURNING id")

	row := query.RunWith(p.db).QueryRowContext(ctx)
	var id int
	if err := row.Scan(&id); err != nil {
		return err
	}

	pet.Id = id
	return nil
}
