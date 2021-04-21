package repository

import (
	"context"
	"pikachu/pkg/domain"
)

type AddressRepository interface {
	CreateUserAddress(ctx context.Context, userId string, address *domain.Address) (error, *domain.Address)
	UpdateUserAddress(ctx context.Context, userId string, addressId string, address *domain.Address) (error, *domain.Address)
	ListUserAddresses(ctx context.Context, userId string) (error, []*domain.Address)
}
