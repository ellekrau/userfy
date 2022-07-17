package namepatternenum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		namePattern NamePatternEnum
		stringValue string
	}{
		{
			namePattern: LowerCase,
			stringValue: "lowercase",
		},
		{
			namePattern: UpperCase,
			stringValue: "uppercase",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.stringValue, tc.namePattern.ToString())
	}
}

func TestIsNamePatternEnumValue(t *testing.T) {
	testCases := []struct {
		input   string
		isValid bool
	}{
		{input: "", isValid: false},
		{input: "title", isValid: false},
		{input: "lowercase", isValid: true},
		{input: "uppercase", isValid: true},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.isValid, IsNamePatternEnumValue(tc.input))
	}
}
