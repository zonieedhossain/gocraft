package usecase

import (
	"github.com/testuser/myapp-test/internal/domain"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) GetAllUsers() ([]domain.User, error) {
	return u.repo.GetAll()
}

func (u *UserUsecase) GetUserByID(id uint) (domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.repo.Create(user)
}
