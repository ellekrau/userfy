package namepatternenum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNamePatternEnumValue(t *testing.T) {
	testCases := []struct {
		Input   string
		IsValid bool
	}{
		{Input: "a", IsValid: false},
		{Input: "abc", IsValid: false},
		{Input: "lowercase", IsValid: true},
		{Input: "uppercase", IsValid: true},
		{Input: "title", IsValid: true},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.IsValid, IsNamePatternEnumValue(tc.Input))
	}
}
