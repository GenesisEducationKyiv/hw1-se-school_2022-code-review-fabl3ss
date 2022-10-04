package usecases

import "genesis_test_case/src/pkg/domain/models"

type SubscriptionUsecase interface {
	Subscribe(recipient *models.Recipient) error
}
