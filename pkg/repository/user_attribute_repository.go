package repository

import (
	"context"
	"pikachu/pkg/domain"
)

type UserAttributeRepository interface {
	CreateUserAttribute(ctx context.Context, userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute)
	UpdateUserAttribute(ctx context.Context, userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute)
	ListUserAttributes(ctx context.Context, userId string) (error, []domain.UserAttribute)
	GetUserAttributeByKey(ctx context.Context, userId string, attributeKey string) (error, *domain.UserAttribute)
}
