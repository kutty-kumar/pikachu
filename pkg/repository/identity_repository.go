package repository

import (
	"context"
	"pikachu/pkg/domain"
)

type IdentityRepository interface {
	CreateIdentity(ctx context.Context, userId string, identity *domain.Identity) (error, *domain.Identity)
	UpdateIdentity(ctx context.Context, userId string, identityId string, identity *domain.Identity) (error, *domain.Identity)
	ListIdentities(ctx context.Context, userId string) (error, []domain.Identity)
}
