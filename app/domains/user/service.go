package user

import (
	"fmt"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/types"
	"github.com/mrrizkin/boot/third-party/hashing"
)

type Service struct {
	repo    *Repo
	hashing hashing.Hashing
}

func NewService(repo *Repo, hashing hashing.Hashing) *Service {
	return &Service{repo, hashing}
}

func (s *Service) Create(user *models.User) (*models.User, error) {
	if user.Password == nil {
		return nil, fmt.Errorf("password is required")
	}

	hash, err := s.hashing.GenerateHash(*user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = &hash
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) FindAll(pagination types.Pagination) (*PaginatedUser, error) {
	users, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	usersCount, err := s.repo.FindAllCount()
	if err != nil {
		return nil, err
	}

	return &PaginatedUser{
		Result: users,
		Total:  int(usersCount),
	}, nil
}

func (s *Service) FindByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Update(id uint, user *models.User) (*models.User, error) {
	userExist, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user.Password != nil {
		if *user.Password != "" {
			hash, err := s.hashing.GenerateHash(*user.Password)
			if err != nil {
				return nil, err
			}

			userExist.Password = &hash
		}
	}

	userExist.Name = user.Name
	userExist.Email = user.Email
	userExist.Username = user.Username
	userExist.RoleID = user.RoleID

	err = s.repo.Update(userExist)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Delete(userLogin *models.User, id uint) error {
	if userLogin.ID == uint(id) {
		return fmt.Errorf("cannot delete yourself")
	}

	return s.repo.Delete(id)
}
