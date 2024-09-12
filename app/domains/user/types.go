package user

import (
	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/database"
	"github.com/mrrizkin/boot/third-party/hashing"
)

type Repo struct {
	db *database.Database
}

type Service struct {
	repo    *Repo
	hashing hashing.Hashing
}

type PaginatedUser struct {
	Result []models.User
	Total  int
}
