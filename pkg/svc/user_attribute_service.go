package svc

import (
	"context"
	"github.com/kutty-kumar/charminder/pkg"
	"pikachu/pkg/domain"
	"pikachu/pkg/repository"
)

type UserAttributeService struct {
	pkg.BaseSvc
	UserAttributeRepo repository.UserAttributeRepository
}

func NewUserAttributeService(baseSvc pkg.BaseSvc, userAttributeRepository repository.UserAttributeRepository) UserAttributeService {
	return UserAttributeService{
		baseSvc,
		userAttributeRepository,
	}
}

func (uas *UserAttributeService) UpdateUserAttribute(ctx context.Context, userId string, other *domain.UserAttribute) (error, *domain.UserAttribute) {
	return uas.UserAttributeRepo.UpdateUserAttribute(ctx, userId, other)
}

func (uas *UserAttributeService) GetUserAttributesByKey(ctx context.Context, userId string, attributeKey string) (error, *domain.UserAttribute) {
	return uas.UserAttributeRepo.GetUserAttributeByKey(ctx, userId, attributeKey)
}

func (uas *UserAttributeService) ListUserAttributes(ctx context.Context, userId string) (error, []domain.UserAttribute) {
	return uas.UserAttributeRepo.ListUserAttributes(ctx, userId)
}
