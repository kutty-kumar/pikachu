package repository

import (
	"context"
	"errors"
	"github.com/kutty-kumar/charminder/pkg"
	"pikachu/pkg/domain"
)

type UserAttributeGORMRepository struct {
	pkg.BaseDao
}

func (ir *UserAttributeGORMRepository) GetUserAttribute(ctx context.Context, userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute) {
	nUserAttribute := &domain.UserAttribute{}
	if err := ir.GetDb().WithContext(ctx).Model(nUserAttribute).Where("user_id = ? AND attribute_key = ?", userId, userAttribute.AttributeKey).Find(nUserAttribute).Error; err != nil {
		return err, nil
	}
	return nil, interface{}(nUserAttribute).(*domain.UserAttribute)
}

func (ir *UserAttributeGORMRepository) CreateUserAttribute(ctx context.Context, userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute) {
	err, eUserAttribute := ir.GetUserAttribute(ctx, userId, userAttribute)
	if (err != nil && eUserAttribute == nil) || (err == nil && eUserAttribute != nil) {
		return errors.New("either user attribute exists or user doesn't exist"), nil
	}
	userAttribute.UserID = userId
	err, cUserAttribute := ir.Create(ctx, userAttribute)
	if err != nil {
		return err, nil
	}
	return nil, cUserAttribute.(*domain.UserAttribute)
}

func (ir *UserAttributeGORMRepository) UpdateUserAttribute(ctx context.Context, userId string, userAttribute *domain.UserAttribute) (error, *domain.UserAttribute) {
	err, eUserAttribute := ir.GetUserAttribute(ctx, userId, userAttribute)
	if err != nil || eUserAttribute == nil {
		return errors.New("either user attribute doesn't exists or user doesn't exist"), nil
	}
	eUserAttribute.Merge(userAttribute)
	err, uUserAttribute := ir.Update(ctx, eUserAttribute.ExternalId, eUserAttribute)
	if err != nil {
		return err, nil
	}
	return nil, interface{}(uUserAttribute).(*domain.UserAttribute)
}

func (ir *UserAttributeGORMRepository) ListUserAttributes(ctx context.Context, userId string) (error, []domain.UserAttribute) {
	user := domain.User{}
	user.ExternalId = userId
	userAttributes := make([]domain.UserAttribute, 0)
	if err := ir.GetDb().WithContext(ctx).Model(&user).Where("external_id = ?", userId).Find(&user).Error; err != nil {
		return err, nil
	}
	if err := ir.GetDb().WithContext(ctx).Model(&user).Association("UserAttributes").Find(&userAttributes); err != nil {
		return err, nil
	}
	return nil, userAttributes
}

func (ir *UserAttributeGORMRepository) GetUserAttributeByKey(ctx context.Context, userId string, attributeKey string) (error, *domain.UserAttribute) {
	userAttr := domain.UserAttribute{}
	if err := ir.GetDb().WithContext(ctx).Model(userAttr).Where("user_id = ? AND attribute_key = ?", userId, attributeKey).Find(userAttr).Error; err != nil {
		return err, nil
	}
	return nil, &userAttr
}

func NewUserAttributeGormRepository(dao pkg.BaseDao) UserAttributeGORMRepository {
	return UserAttributeGORMRepository{
		dao,
	}
}
