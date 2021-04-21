package repository

import (
	"context"
	"github.com/kutty-kumar/charminder/pkg"
	"pikachu/pkg/domain"
)

type AddressGORMRepository struct {
	pkg.BaseRepository
}

func (ar *AddressGORMRepository) CreateUserAddress(ctx context.Context, userId string, address *domain.Address) (error, *domain.Address) {
	err, user := ar.GetByExternalId(ctx, userId)
	if err != nil {
		return err, nil
	}
	address.UserID = user.GetExternalId()
	err, bAddress := ar.Create(ctx, address)
	if err != nil {
		return err, nil
	}
	return nil, bAddress.(*domain.Address)
}

func (ar *AddressGORMRepository) UpdateUserAddress(ctx context.Context, userId string, addressId string, address *domain.Address) (error, *domain.Address) {
	err, user := ar.GetByExternalId(ctx, userId)
	if err != nil {
		return err, nil
	}
	eAddress := domain.Address{}
	if err := ar.GetDb().Model(&eAddress).Where("external_id = ? AND user_id = ?", addressId, user.GetId()).Find(&eAddress).Error; err != nil {
		return err, nil
	}
	eAddress.Merge(address)
	err, uAddress := ar.Update(ctx, addressId, &eAddress)
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uAddress).(*domain.Address)
}

func (ar *AddressGORMRepository) ListUserAddresses(ctx context.Context, userId string) (error, []*domain.Address) {
	addresses := make([]*domain.Address, 0)
	err, user := ar.GetByExternalId(ctx, userId)
	if err != nil {
		return err, nil
	}
	if err := ar.GetDb().WithContext(ctx).Model(user).Association("Address").Find(&addresses); err != nil {
		return err, nil
	}
	return nil, addresses
}
