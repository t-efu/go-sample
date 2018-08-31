package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sample/application/entity/model"
	"github.com/sample/application/repository/mock"
	"github.com/sample/application/usecase"
	"github.com/stretchr/testify/assert"
)

var (
	exampleUser = model.User{
		ID:   1,
		Name: "test_user",
	}
	exampleUser2 = model.User{
		ID:   2,
		Name: "test_user_2",
	}
)

func TestUserRepositoryGet(t *testing.T) {
	cases := []struct {
		expectObject *model.User
		expectError  string
		ctx          context.Context
		Mock         func(*mock.MockUserRepository)
		id           uint64
	}{
		{
			expectError: "failed get user: something",
			ctx:         context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("something"))
			},
			id: exampleUser.ID,
		},
		{
			expectObject: &exampleUser,
			ctx:          context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&exampleUser, nil)
			},
			id: exampleUser.ID,
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockUserRepository(ctrl)
		if c.Mock != nil {
			c.Mock(repo)
		}
		usecase := usecase.NewUserUsecase(repo)
		actual, err := usecase.Get(c.ctx, c.id)
		if actual != nil {
			assert.Equal(t, c.expectObject, actual)
			assert.Nil(t, err)
		}
		if err != nil {
			assert.EqualError(t, err, c.expectError)
		}
	}
}

func TestUserRepositoryFind(t *testing.T) {
	cases := []struct {
		expectObjects []model.User
		expectError   string
		ctx           context.Context
		Mock          func(*mock.MockUserRepository)
	}{
		{
			expectError: "failed find users: something",
			ctx:         context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Find(gomock.Any()).Return(nil, errors.New("something"))
			},
		},
		{
			expectObjects: []model.User{exampleUser, exampleUser2},
			ctx:           context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Find(gomock.Any()).Return([]model.User{exampleUser, exampleUser2}, nil)
			},
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockUserRepository(ctrl)
		if c.Mock != nil {
			c.Mock(repo)
		}
		usecase := usecase.NewUserUsecase(repo)
		actual, err := usecase.Find(c.ctx)
		if actual != nil {
			assert.Equal(t, c.expectObjects, actual)
			assert.Nil(t, err)
		}
		if err != nil {
			assert.EqualError(t, err, c.expectError)
		}
	}
}

func TestUserRepositoryCreate(t *testing.T) {
	cases := []struct {
		expectObject *model.User
		expectError  string
		ctx          context.Context
		Mock         func(*mock.MockUserRepository)
		user         *model.User
	}{
		{
			expectError: "failed create user: something",
			ctx:         context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("something"))
			},
			user: &model.User{Name: exampleUser.Name},
		},
		{
			expectObject: &exampleUser,
			ctx:          context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&exampleUser, nil)
			},
			user: &model.User{Name: exampleUser.Name},
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockUserRepository(ctrl)
		if c.Mock != nil {
			c.Mock(repo)
		}
		usecase := usecase.NewUserUsecase(repo)
		actual, err := usecase.Create(c.ctx, c.user)
		if actual != nil {
			assert.Equal(t, c.expectObject, actual)
			assert.Nil(t, err)
		}
		if err != nil {
			assert.EqualError(t, err, c.expectError)
		}
	}
}

func TestUserRepositoryUpdate(t *testing.T) {
	cases := []struct {
		expectError string
		ctx         context.Context
		Mock        func(*mock.MockUserRepository)
		user        *model.User
	}{
		{
			expectError: "failed update user: something",
			ctx:         context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("something"))
			},
			user: &model.User{Name: exampleUser.Name},
		},
		{
			ctx: context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
			},
			user: &model.User{Name: exampleUser.Name},
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockUserRepository(ctrl)
		if c.Mock != nil {
			c.Mock(repo)
		}
		usecase := usecase.NewUserUsecase(repo)
		err := usecase.Update(c.ctx, c.user)
		if err != nil {
			assert.EqualError(t, err, c.expectError)
		}
	}
}

func TestUserRepositoryDelete(t *testing.T) {
	cases := []struct {
		expectError string
		ctx         context.Context
		Mock        func(*mock.MockUserRepository)
		id          uint64
	}{
		{
			expectError: "failed delete user: something",
			ctx:         context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("something"))
			},
			id: exampleUser.ID,
		},
		{
			ctx: context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			id: exampleUser.ID,
		},
	}

	for _, c := range cases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock.NewMockUserRepository(ctrl)
		if c.Mock != nil {
			c.Mock(repo)
		}
		usecase := usecase.NewUserUsecase(repo)
		err := usecase.Delete(c.ctx, c.id)
		if err != nil {
			assert.EqualError(t, err, c.expectError)
		}
	}
}
