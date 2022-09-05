package usecase

import (
	"errors"
	"genesis_test_case/src/pkg/domain"
	mocks "genesis_test_case/src/pkg/domain/mocks"
	"genesis_test_case/src/pkg/repository"
	myerr "genesis_test_case/src/pkg/types/errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestSubscribe(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockMailingRepo := mocks.NewMockMailingRepository(ctl)
	repoCollection := &repository.Repositories{
		Mailing: mockMailingRepo,
	}
	mailingUcase := NewMailingUsecase(repoCollection)
	recipient := &domain.Recipient{
		Email: "test@test.com",
	}
	mockMailingResp := []string{
		"example@example.com",
		"xyz@xyz.com",
	}

	mockMailingRepo.EXPECT().GetSubscribed().Return(mockMailingResp, nil)
	mockMailingRepo.EXPECT().InsertNewEmail(mockMailingResp, recipient.Email).Return(nil)
	err := mailingUcase.Subscribe(recipient)
	require.NoError(t, err)
}

func TestSubscribeError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockMailingRepo := mocks.NewMockMailingRepository(ctl)
	repoCollection := &repository.Repositories{Mailing: mockMailingRepo}
	mailingUcase := NewMailingUsecase(repoCollection)
	recipient := &domain.Recipient{
		Email: "test@test.com",
	}
	mockMailingResp := []string{
		"example@example.com",
		"xyz@xyz.com",
	}

	err := mailingUcase.Subscribe(nil)
	require.EqualError(t, err, myerr.ErrNoDataProvided.Error())

	mockMailingRepo.EXPECT().GetSubscribed().Return(nil, errors.New("error!"))
	err = mailingUcase.Subscribe(recipient)
	require.Error(t, err)

	mockMailingRepo.EXPECT().GetSubscribed().Return(mockMailingResp, nil)
	mockMailingRepo.EXPECT().InsertNewEmail(mockMailingResp, recipient.Email).Return(myerr.ErrAlreadyExists)
	err = mailingUcase.Subscribe(recipient)
	require.EqualError(t, err, myerr.ErrAlreadyExists.Error())
}

func TestSendCurrencyRate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockMailingRepo := mocks.NewMockMailingRepository(ctl)
	mockCryptoRepo := mocks.NewMockCryptoRepository(ctl)
	mockImageRepo := mocks.NewMockImageRepository(ctl)
	repoCollection := &repository.Repositories{
		Mailing: mockMailingRepo,
		Crypto:  mockCryptoRepo,
		Image:   mockImageRepo,
	}
	mailingUcase := NewMailingUsecase(repoCollection)
	mockMailingBodyResp := "example"
	mockMailingSendResp := []string{"example@example.com"}
	mockCryptoChartResp := []float64{0.0, 0.1, 0.2}
	mockCryptoRateResp := &domain.CurrencyRate{}
	mockImageResp := "http://example.com/example"

	mockCryptoRepo.EXPECT().GetWeekChart().Return(mockCryptoChartResp, nil)
	mockCryptoRepo.EXPECT().GetCurrencyRate("", "").Return(mockCryptoRateResp, nil)
	mockImageRepo.EXPECT().GetCryptoBannerUrl(mockCryptoChartResp, mockCryptoRateResp).Return(mockImageResp, nil)
	mockMailingRepo.EXPECT().GetMessageBody(mockImageResp).Return(mockMailingBodyResp, nil)
	mockMailingRepo.EXPECT().SendToSubscribed(mockMailingBodyResp).Return(mockMailingSendResp, nil)
	unsent, err := mailingUcase.SendCurrencyRate()
	require.NoError(t, err)
	require.Equal(t, unsent, mockMailingSendResp)
}

func TestSendCurrencyRateError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockMailingRepo := mocks.NewMockMailingRepository(ctl)
	mockCryptoRepo := mocks.NewMockCryptoRepository(ctl)
	u := &repository.Repositories{
		Mailing: mockMailingRepo,
		Crypto:  mockCryptoRepo,
	}
	mailingUcase := NewMailingUsecase(u)

	mockCryptoRepo.EXPECT().GetWeekChart().Return(nil, errors.New("error!"))
	_, err := mailingUcase.SendCurrencyRate()
	require.Error(t, err)
}
