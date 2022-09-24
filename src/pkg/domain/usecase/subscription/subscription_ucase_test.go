package usecase

import (
	"genesis_test_case/src/pkg/domain/models"
	mocks "genesis_test_case/src/pkg/domain/usecase/mocks"
	myerr "genesis_test_case/src/pkg/types/errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestSubscribe(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mocks.NewMockEmailStorage(ctl)
	subscriptionUsecase := NewSubscriptionUsecase(
		mockStorage,
	)
	recipient := &models.Recipient{
		Email: "test@test.com",
	}

	mockStorage.EXPECT().AddEmail(recipient.Email).Return(nil)
	err := subscriptionUsecase.Subscribe(recipient)
	require.NoError(t, err)
}

func TestSubscribeError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockStorage := mocks.NewMockEmailStorage(ctl)
	subscriptionUsecase := NewSubscriptionUsecase(
		mockStorage,
	)
	recipient := &models.Recipient{
		Email: "test@test.com",
	}

	mockStorage.EXPECT().AddEmail(recipient.Email).Return(myerr.ErrAlreadyExists)
	err := subscriptionUsecase.Subscribe(recipient)
	require.EqualError(t, err, myerr.ErrAlreadyExists.Error())
}
