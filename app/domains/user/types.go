package user

import (
	"github.com/mrrizkin/boot/app/models"
)

type PaginatedUser struct {
	Result []models.User
	Total  int
}
