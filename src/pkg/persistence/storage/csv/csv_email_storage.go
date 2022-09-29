package storage

import (
	"encoding/csv"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	myerr "genesis_test_case/src/pkg/types/errors"
	"genesis_test_case/src/pkg/types/filemodes"
	"genesis_test_case/src/pkg/utils"
	"os"
)

type csvEmailStorage struct {
	csvPath string
}

func NewCsvEmaiStorage(path string) application.EmailStorage {
	return &csvEmailStorage{
		csvPath: path,
	}
}

func (c *csvEmailStorage) AddEmail(toInsert models.EmailAddress) error {
	emails, err := c.GetAllEmails()
	if err != nil {
		return err
	}

	cmp := func(a models.EmailAddress, b models.EmailAddress) int {
		if a.Address == b.Address {
			return 0
		}
		if a.Address < b.Address {
			return -1
		}
		return 1
	}
	sorted, err := utils.InsertToSorted(emails, toInsert, cmp)
	if err != nil {
		return err
	}
	err = c.insertEmailsToCsvFile(c.csvPath, sorted)
	if err != nil {
		return err
	}

	return nil
}

func (c *csvEmailStorage) GetAllEmails() ([]models.EmailAddress, error) {
	subscribed, err := utils.ReadAllFromCsv(c.csvPath)
	if err != nil {
		return nil, err
	}

	var emails []models.EmailAddress
	for _, s := range subscribed {
		emails = append(emails, models.EmailAddress{Address: s})
	}

	return emails, nil
}

func (c *csvEmailStorage) insertEmailsToCsvFile(path string, data []models.EmailAddress) error {
	if path == "" || len(data) < 1 {
		return myerr.ErrInvalidInput
	}
	fileMode := os.ModeDir | (filemodes.OS_USER_RW | filemodes.OS_ALL_R)
	f, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		fileMode,
	)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	w := csv.NewWriter(f)
	defer w.Flush()

	for _, v := range data {
		err := w.Write([]string{v.Address})
		if err != nil {
			return err
		}
	}

	return nil
}
