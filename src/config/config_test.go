package config

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCellphonePattern(t *testing.T) {
	testCases := []struct {
		TestCase         string
		CellphonePattern string
		Error            error
	}{
		{
			TestCase:         "CellphonePattern only with letters",
			CellphonePattern: "none",
			Error:            errors.New(errInvalidCellphonePattern),
		},
		{
			TestCase:         "CellphonePattern without enough Xs (6)",
			CellphonePattern: "XXXXXX",
			Error:            errors.New(errInvalidCellphonePattern),
		},
		{
			TestCase:         "CellphonePattern with too many 0s (14)",
			CellphonePattern: "XXXXXXXXXXXXXX",
			Error:            errors.New(errInvalidCellphonePattern),
		},
		{
			TestCase:         "CellphonePattern with empty value",
			CellphonePattern: "",
			Error:            nil,
		},
		{
			TestCase:         "CellphonePattern with empty value",
			CellphonePattern: "XXXXXXXXXXXXX",
			Error:            nil,
		},
	}

	for _, tc := range testCases {
		UserData.CellphonePattern = tc.CellphonePattern
		assert.Equal(t, tc.Error, validateCellphonePattern(), tc.TestCase)
	}
}
