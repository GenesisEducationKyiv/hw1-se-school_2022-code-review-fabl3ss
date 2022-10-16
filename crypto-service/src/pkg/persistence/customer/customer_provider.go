package customer

import (
	"bytes"
	"encoding/json"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"net/http"
)

func NewCustomerProvider(ordersURL string) application.CustomerProvider {
	return &customerProvider{
		ordersServiceURL: ordersURL,
	}
}

type customerProvider struct {
	ordersServiceURL string
}

func (c *customerProvider) CreateCustomer(recipient *models.Recipient) error {
	body, err := c.createCustomerRequestBody(recipient)
	if err != nil {
		return err
	}
	req, err := c.createCustomerRequest(body)
	if err != nil {
		return err
	}

	return c.createCustomerMakeRequest(req)
}

func (c *customerProvider) createCustomerRequest(body []byte) (*http.Request, error) {
	req, err := http.NewRequest(
		"POST",
		c.ordersServiceURL+"/register-customer",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *customerProvider) createCustomerRequestBody(recipient *models.Recipient) ([]byte, error) {
	reqBody := struct {
		EmailAddress string `json:"emailAddress"`
	}{EmailAddress: recipient.Email.Address}

	return json.Marshal(reqBody)
}

func (c *customerProvider) createCustomerMakeRequest(req *http.Request) error {
	client := &http.Client{}

	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err = r.Body.Close()
	}()

	return err
}
