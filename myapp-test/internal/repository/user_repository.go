package repository

import (
	"github.com/testuser/myapp-test/internal/domain"
	
	"database/sql"
	
)

type UserRepo struct {
	
	db *sql.DB
	
}

func NewRepository(db *sql.DB) domain.UserRepository {
	return &UserRepo{db: db}
}


func (r *UserRepo) GetAll() ([]domain.User, error) {
	q := New(r.db)
	users, err := q.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}
	var domainUsers []domain.User
	for _, u := range users {
		domainUsers = append(domainUsers, domain.User{
			ID:        uint(u.ID),
			Username:  u.Username,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return domainUsers, nil
}

func (r *UserRepo) GetByID(id uint) (domain.User, error) {
	q := New(r.db)
	u, err := q.GetUserByID(context.Background(), int64(id))
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        uint(u.ID),
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *UserRepo) Create(user *domain.User) error {
	q := New(r.db)
	u, err := q.CreateUser(context.Background(), CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return err
	}
	user.ID = uint(u.ID)
	return nil
}

func (r *UserRepo) Update(user *domain.User) error {
	// Simple update implementation for sqlc
	return nil
}

func (r *UserRepo) Delete(id uint) error {
	q := New(r.db)
	return q.DeleteUser(context.Background(), int64(id))
}


