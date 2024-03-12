package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"studentgit.kata.academy/xp/PetStore/internal/store/entities"
)

type StoreRepository interface {
	Create(ctx context.Context, store *entities.Store) error
	Get(ctx context.Context, id int) (*entities.Store, error)
	Delete(ctx context.Context, id int) error
}

type storeRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(db *sql.DB) StoreRepository {
	return &storeRepository{
		db:         db,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (s *storeRepository) Create(ctx context.Context, store *entities.Store) error {
	query := s.sqlBuilder.Insert("orders").Columns("pet_id", "ship_date", "status", "complete").
		Values(store.PetId, store.ShipDate, store.Status, store.Complete).
		Suffix("RETURNING id")

	row := query.RunWith(s.db).QueryRowContext(ctx)
	return row.Scan(&store.Id)
}

func (s *storeRepository) Get(ctx context.Context, id int) (*entities.Store, error) {
	query := s.sqlBuilder.Select("id", "pet_id", "ship_date", "status", "complete").
		From("orders").
		Where(sq.Eq{"id": id})

	row := query.RunWith(s.db).QueryRowContext(ctx)
	var store entities.Store
	if err := row.Scan(&store.Id, &store.PetId, &store.ShipDate, &store.Status, &store.Complete); err != nil {
		return nil, err
	}

	return &store, nil
}

func (s *storeRepository) Delete(ctx context.Context, id int) error {
	query := s.sqlBuilder.Delete("orders").Where(sq.Eq{"id": id})

	res, err := query.RunWith(s.db).ExecContext(ctx)
	if err != nil {
		return err
	}

	if _, err := res.RowsAffected(); err != nil {
		return err
	}
	return nil
}
