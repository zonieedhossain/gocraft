package repository

import (
	"github.com/testuser/myapp-test/internal/domain"
	
	"gorm.io/gorm"
	
)

type UserRepo struct {
	
	db *gorm.DB
	
}

func NewRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepo{db: db}
}


func (r *UserRepo) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepo) GetByID(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}


