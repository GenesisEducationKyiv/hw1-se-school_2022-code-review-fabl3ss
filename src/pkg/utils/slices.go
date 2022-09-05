package utils

import (
	"genesis_test_case/src/pkg/types/errors"
	"sort"
)

func InsertToSorted(s []string, toInsert string) ([]string, error) {
	if s == nil {
		return nil, errors.ErrInvalidInput
	}
	index := sort.SearchStrings(s, toInsert)
	if index != len(s) {
		if s[index] == toInsert {
			return nil, errors.ErrAlreadyExists
		}
	}

	s = append(s, "")
	copy(s[index+1:], s[index:])
	s[index] = toInsert

	return s, nil
}
