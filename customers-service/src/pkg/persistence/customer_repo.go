package persistence

import (
	"customers/src/pkg/application"
	"customers/src/pkg/domain/models"

	"gorm.io/gorm"
)

func NewCustomerRepository(db *gorm.DB) application.CustomerRepository {
	return &customerRepo{
		db: db,
	}
}

type customerRepo struct {
	db *gorm.DB
}

func (c customerRepo) CreateCustomer(customer *models.Customer) error {
	return c.db.Create(customer).Error
}

func (c customerRepo) UpdateCustomerStatus(transactionId, status string) error {
	return c.db.Model(&models.Customer{}).
		Where("id_transaction = ?", transactionId).
		Update("status", status).
		Limit(1).
		Error
}
