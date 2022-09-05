package utils

import (
	"encoding/csv"
	"genesis_test_case/src/pkg/types/filemodes"
	"io"
	"os"
)

func WriteToCsv(path string, data []string) error {
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
		err := w.Write([]string{v})
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadAllFromCsv(path string) ([]string, error) {
	fileMode := os.ModeDir | (filemodes.OS_USER_RW | filemodes.OS_ALL_R)
	f, err := os.OpenFile(
		path,
		os.O_RDONLY|os.O_CREATE,
		fileMode,
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	content, err := csvFileToSlice(f)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func csvFileToSlice(f *os.File) ([]string, error) {
	var data []string
	csvReader := csv.NewReader(f)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, record[0])
	}

	return data, nil
}
