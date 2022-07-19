package nameformatter

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatName(t *testing.T) {
	testCases := []struct {
		testCase       string
		namePattern    string
		expectedResult string
	}{
		{
			testCase:       "Lower case",
			namePattern:    "lowercase",
			expectedResult: "testing the function",
		},
		{
			testCase:       "Upper case",
			namePattern:    "uppercase",
			expectedResult: "TESTING THE FUNCTION",
		},
		{
			testCase:       "None",
			namePattern:    "",
			expectedResult: "Testing the function",
		},
	}

	for _, tc := range testCases {
		config.OldUserData.NamePattern = tc.namePattern
		assert.Equal(t, tc.expectedResult, FormatName("Testing the function"), tc.testCase)
	}
}
