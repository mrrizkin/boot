package user

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/database"
)

type UserRepo struct {
	db *database.Database
}

type UserRepoDeps struct {
	fx.In

	Database *database.Database
}

func NewUserRepo(p UserRepoDeps) (*UserRepo, error) {
	return &UserRepo{
		db: p.Database,
	}, nil
}

func (r *UserRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepo) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
