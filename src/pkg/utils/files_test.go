package utils

import (
	"genesis_test_case/src/config"
	myerr "genesis_test_case/src/pkg/types/errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestWriteToCsv(t *testing.T) {
	testPath := "../../platform/csv/test.csv"
	testData := []string{"example1", "example2"}
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	err := WriteToCsv(testPath, testData)
	require.NoError(t, err)
	err = os.Remove(testPath)
}

func TestWriteToCsvError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	err := WriteToCsv("", []string{"test"})
	require.EqualError(t, err, myerr.ErrInvalidInput.Error())
	err = WriteToCsv("test_path", nil)
	require.EqualError(t, err, myerr.ErrInvalidInput.Error())
}

func TestReadAllFromCsv(t *testing.T) {
	testPath := "../../platform/csv/test.csv"
	testData := []string{"example1", "example2"}
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	err := WriteToCsv(testPath, testData)
	require.NoError(t, err)
	data, err := ReadAllFromCsv(testPath)
	require.NoError(t, err)
	require.Equal(t, testData, data)
}

func TestStorageIsEmpty(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	cfg := config.Get()

	data, err := ReadAllFromCsv(cfg.StorageFile)
	require.Error(t, err)
	require.Empty(t, data)
}
