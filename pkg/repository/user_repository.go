package repository

import (
	"context"
	"github.com/kutty-kumar/charminder/pkg"
	"pikachu/pkg/domain"
)

type UserRepository interface {
	pkg.BaseRepository
	GetUserByEmailPassword(ctx context.Context, email string, password string) (*domain.User, error)
}
