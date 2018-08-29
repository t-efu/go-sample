package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sample/application/entity/model"
)

// UserRepository user repository
type UserRepository interface {
	Get(ctx context.Context, id uint64) (*model.User, error)
	Find(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}

// NewUserRepository new user repository
func NewUserRepository(Conn *gorm.DB) UserRepository {
	return &userRepository{Conn}
}

type userRepository struct {
	Conn *gorm.DB
}

func (r *userRepository) Get(ctx context.Context, id uint64) (*model.User, error) {
	user := &model.User{}
	err := r.Conn.Where("id = ?", id).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed get user")
	}
	return user, nil
}

func (r *userRepository) Find(ctx context.Context) ([]model.User, error) {
	users := make([]model.User, 5)
	err := r.Conn.Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed find users")
	}
	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.Conn.Create(user).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed create user")
	}
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	err := r.Conn.Model(model.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uint64) error {
	err := r.Conn.Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return errors.Wrap(err, "failed delete user")
	}
	return nil
}
