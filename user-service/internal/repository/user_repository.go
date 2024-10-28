package repository

import (
	"context"
	"database/sql"
	"errors"
	"user-service/internal/domain/model"
	"user-service/internal/ports"
)

type UserRepository struct {
	db *sql.DB
}

type UserGetter interface {
	User() ports.UserRepository
}

// Implement the User() method for UserGetter interface
func (u *UserRepository) User() ports.UserRepository {
	return u
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// create user in database
	query := "INSERT INTO users (name,email,address,password) VALUES($1,$2,$3,$4) RETURNING id"
	err := u.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Address, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func (u *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	// get all users from database
	query := "SELECT * FROM users"
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err

	}
	return users, nil
}
func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	// Include password in the query
	query := "SELECT id, name, email, address, password FROM users WHERE email = $1"
	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user not found
		}
		return nil, err // other errors
	}
	return &user, nil
}

func (u *UserRepository) GetUserById(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	// Include password in the query
	query := "SELECT id, name, email, address, password FROM users WHERE id = $1"
	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user not found
		}
		return nil, err // other errors
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// Update user in the database
	query := `UPDATE users SET name= $1,email= $2,address= $3 WHERE id= $4`
	result, err := u.db.ExecContext(ctx, query, user.Name, user.Email, user.Address, user.ID)
	if err != nil {
		return nil, err
	}
	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err

	}
	if rowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}
