package utils

import (
	"genesis_test_case/src/pkg/types/errors"

	"golang.org/x/exp/slices"
)

func InsertToSorted[T any](array []T, value T, cmp func(T, T) int) ([]T, error) {
	pos, isFound := slices.BinarySearchFunc(array, value, cmp)
	if isFound {
		return nil, errors.ErrAlreadyExists
	}

	slices.Insert(array, pos, value)

	return array, nil
}
