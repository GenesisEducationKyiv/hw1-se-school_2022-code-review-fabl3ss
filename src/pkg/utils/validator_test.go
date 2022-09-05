package utils

import (
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/types/errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateStruct(t *testing.T) {
	input := &domain.Recipient{
		Email: "test@test.com",
	}
	_, err := ValidateStruct(input)
	require.NoError(t, err)
}

func TestValidateStructError(t *testing.T) {
	cases := []struct {
		testName    string
		input       interface{}
		expectedErr error
	}{
		{
			testName:    "no_body",
			input:       nil,
			expectedErr: errors.ErrNoDataProvided,
		},
		{
			testName: "bad_email",
			input: &domain.Recipient{
				Email: "testemail123",
			},
			expectedErr: errors.ErrValidationFailed,
		},
	}

	for _, tcase := range cases {
		t.Run(tcase.testName, func(t *testing.T) {
			msg, err := ValidateStruct(tcase.input)
			require.EqualError(t, err, tcase.expectedErr.Error())
			require.NotEmpty(t, msg)
		})
	}
}
