package user

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/model"
)

type UserService struct {
	repo *UserRepo
}

type UserServiceDeps struct {
	fx.In

	Repo *UserRepo
}

func NewUserService(p UserServiceDeps) (*UserService, error) {
	return &UserService{
		repo: p.Repo,
	}, nil
}

func (s *UserService) Create(user *model.User) (*model.User, error) {
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindAll() ([]model.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) FindByID(id uint) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(id uint, user *model.User) (*model.User, error) {
	var err error

	_, err = s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}
