package repository

import (
	"context"
	"errors"
	"github.com/kutty-kumar/charminder/pkg"
	"pikachu/pkg/domain"
)

type UserGormRepository struct {
	pkg.BaseDao
}

func NewUserGormRepository(dao pkg.BaseDao) UserGormRepository {
	return UserGormRepository{
		dao,
	}
}

func (u *UserGormRepository) Create(ctx context.Context, user *domain.User) (error, *domain.User) {
	return u.Create(ctx, user)
}

func (u *UserGormRepository) Update(ctx context.Context, id string, user *domain.User) (error, *domain.User) {
	return u.Update(ctx, id, user)
}

func (u *UserGormRepository) FindByExternalId(ctx context.Context, id string) (error, *domain.User) {
	err, base := u.GetByExternalId(ctx, id)
	if base != nil {
		return err, base.(*domain.User)
	}
	return errors.New("not found"), nil
}

func (u *UserGormRepository) MultiGetByExternalIds(ctx context.Context, ids []string) (error, []domain.User) {
	var userSlice []domain.User
	err, sqlSlice := u.MultiGetByExternalId(ctx, ids)
	if err != nil {
		return err, nil
	}

	for _, sqlRow := range sqlSlice {
		userSlice = append(userSlice, *interface{}(sqlRow).(*domain.User))
	}
	return nil, userSlice
}

func (u *UserGormRepository) handleError(user *domain.User, err error) (error, *domain.User) {
	if err != nil {
		return err, nil
	}
	return nil, user
}
