package utils

import (
	"genesis_test_case/src/pkg/types/errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertToSorted(t *testing.T) {
	testCases := []struct {
		inputSlice  []string
		inputToFind string
	}{
		{
			inputSlice:  []string{"apple", "citrone"},
			inputToFind: "ball",
		},
	}
	cmp := func(a string, b string) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		}
		return 1
	}

	for _, tcase := range testCases {
		s, err := InsertToSorted(tcase.inputSlice, tcase.inputToFind, cmp)
		require.NoError(t, err)
		require.IsIncreasing(t, s)
	}
}

func TestInsertToSortedError(t *testing.T) {
	testCases := []struct {
		name        string
		inputSlice  []string
		inputToFind string
		expErr      error
	}{
		{
			name:        "duplicate",
			inputSlice:  []string{"apple", "ball"},
			inputToFind: "ball",
			expErr:      errors.ErrAlreadyExists,
		},
	}
	cmp := func(a string, b string) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		}
		return 1
	}

	for _, tcase := range testCases {
		t.Run(tcase.name, func(t *testing.T) {
			_, err := InsertToSorted(tcase.inputSlice, tcase.inputToFind, cmp)
			require.EqualError(t, tcase.expErr, err.Error())
		})
	}
}
