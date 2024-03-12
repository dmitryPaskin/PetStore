package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"

	"studentgit.kata.academy/xp/PetStore/internal/user/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	GetIdByUsername(ctx context.Context, username string) (int, error)
	Update(ctx context.Context, username string, user *entities.User) error
	Delete(ctx context.Context, username string) error
	Login(ctx context.Context, userId int) (int, error)
	Logout(ctx context.Context, sessionId int) error
}

type userRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(db *sql.DB) UserRepository {
	return &userRepository{
		db:         db,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (u *userRepository) Create(ctx context.Context, user *entities.User) error {
	query := u.sqlBuilder.Insert("users").
		Columns("username", "first_name", "last_name", "email", "phone", "password", "user_status").
		Values(user.Username, user.FirstName, user.LastName, user.Email, user.Phone, user.Password, user.UserStatus)

	_, err := query.RunWith(u.db).ExecContext(ctx)

	return err
}

func (u *userRepository) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := u.sqlBuilder.Select("id", "username", "first_name", "last_name", "email", "phone", "password", "user_status").
		From("users").Where(sq.Eq{"username": username})

	row := query.RunWith(u.db).QueryRowContext(ctx)

	user := &entities.User{}

	err := row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.Password, &user.UserStatus)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetIdByUsername(ctx context.Context, username string) (int, error) {
	query := u.sqlBuilder.Select("id").
		From("users").Where(sq.Eq{"username": username})

	row := query.RunWith(u.db).QueryRowContext(ctx)

	var id int

	err := row.Scan(id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepository) Update(ctx context.Context, username string, user *entities.User) error {
	query := u.sqlBuilder.Update("users").
		Set("username", user.Username).
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("email", user.Email).
		Set("phone", user.Phone).
		Set("password", user.Password).
		Set("user_status", user.UserStatus).
		Where(sq.Eq{"username": username})

	_, err := query.RunWith(u.db).ExecContext(ctx)
	return err
}

func (u *userRepository) Delete(ctx context.Context, username string) error {
	query := u.sqlBuilder.Delete("users").Where(sq.Eq{"username": username})

	_, err := query.RunWith(u.db).ExecContext(ctx)
	return err
}

func (u *userRepository) Login(ctx context.Context, userId int) (int, error) {
	query := u.sqlBuilder.Insert("sessions").
		Columns("user_id").Values(userId).Suffix("RETURNING id")

	raw := query.RunWith(u.db).QueryRowContext(ctx)
	var sessionId int
	err := raw.Scan(&sessionId)
	return sessionId, err
}

func (u *userRepository) Logout(ctx context.Context, sessionId int) error {
	query := u.sqlBuilder.Delete("sessions").Where(sq.Eq{"id": sessionId})
	_, err := query.RunWith(u.db).ExecContext(ctx)

	return err
}
