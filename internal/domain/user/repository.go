package user

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/model"
	"github.com/mrrizkin/boot/internal/system/database"
)

type UserRepo struct {
	db *database.Database
}

type UserRepoParams struct {
	fx.In

	Database *database.Database
}

func NewUserRepo(p UserRepoParams) (*UserRepo, error) {
	return &UserRepo{
		db: p.Database,
	}, nil
}

func (r *UserRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepo) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepo) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
