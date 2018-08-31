package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sample/application/entity/model"
	"github.com/sample/application/repository"
)

// UserUsecase user usecase
type UserUsecase interface {
	Get(ctx context.Context, id uint64) (*model.User, error)
	Find(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint64) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

// NewUserUsecase new user usecase
func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Get(ctx context.Context, id uint64) (*model.User, error) {
	user, err := u.userRepository.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	return user, nil
}

func (u *userUsecase) Find(ctx context.Context) ([]model.User, error) {
	users, err := u.userRepository.Find(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed find users")
	}
	return users, nil
}

func (u *userUsecase) Create(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed create user")
	}
	return user, nil
}

func (u *userUsecase) Update(ctx context.Context, user *model.User) error {
	err := u.userRepository.Update(ctx, user)
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id uint64) error {
	err := u.userRepository.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed delete user")
	}
	return nil
}
