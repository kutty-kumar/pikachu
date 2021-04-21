package repository

import (
	"context"
	"pikachu/pkg/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (error, *domain.User)
	Update(ctx context.Context, id string, user *domain.User) (error, *domain.User)
	FindByExternalId(ctx context.Context, id string) (error, *domain.User)
	MultiGetByExternalIds(ctx context.Context, ids []string) (error, []domain.User)
}
