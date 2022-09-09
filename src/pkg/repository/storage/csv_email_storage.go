package storage

import (
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/pkg/utils"
)

type csvEmailStorage struct {
	csvPath string
}

func NewCsvEmaiStorage(path string) usecase.EmailStorage {
	return &csvEmailStorage{
		csvPath: path,
	}
}

func (c *csvEmailStorage) AddEmail(toInsert string) error {
	emails, err := c.GetAllEmails()
	if err != nil {
		return err
	}
	sorted, err := utils.InsertToSorted(emails, toInsert)
	if err != nil {
		return err
	}
	err = utils.WriteToCsv(c.csvPath, sorted)
	if err != nil {
		return err
	}

	return nil
}

func (c *csvEmailStorage) GetAllEmails() ([]string, error) {
	subscribed, err := utils.ReadAllFromCsv(c.csvPath)
	if err != nil {
		return nil, err
	}

	return subscribed, nil
}
