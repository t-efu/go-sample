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

func TestUserRepositoryGet(t *testing.T) {
	cases := []struct {
		expectedObject model.User
		expectedError  string
		ctx            context.Context
		Mock           func(*mock.MockUserRepository)
		id             uint64
	}{
		{
			expectedError: "failed get user: something",
			ctx:           context.Background(),
			Mock: func(repo *mock.MockUserRepository) {
				repo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, errors.New("something"))
			},
			id: uint64(1),
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
			assert.Equal(t, c.expectedObject, actual)
			assert.Nil(t, err)
		}
		if err != nil {
			assert.EqualError(t, err, c.expectedError)
		}
	}
}
