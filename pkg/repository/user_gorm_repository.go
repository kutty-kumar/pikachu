package repository

import (
	"context"
	"errors"
	"github.com/kutty-kumar/charminder/pkg"
	"pikachu/pkg/domain"
)

type UserGormRepository struct {
	pkg.BaseRepository
}

func NewUserGormRepository(repo pkg.BaseRepository) UserGormRepository {
	return UserGormRepository{
		repo,
	}
}
func (ugr *UserGormRepository) GetUserByEmailPassword(ctx context.Context, email string, password string) (*domain.User, error) {
	user := domain.User{}
	if err := ugr.GetDb().Model(&user).WithContext(ctx).Joins("INNER JOIN identities ON identities.user_id=users.external_id").Where("identities.identity_type=2 AND identities.identity_value = ? AND users.password = ?", email, password).Find(&user).Error; err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, errors.New("user not found")

	}
	return &user, nil
}
