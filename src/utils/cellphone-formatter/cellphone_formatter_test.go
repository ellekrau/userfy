package cellphoneformatter

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatCellphone(t *testing.T) {
	testCases := []struct {
		Pattern  string
		Expected string
	}{
		{
			Pattern:  "+00 (00) 00000-0000",
			Expected: "+55 (41) 99433-6786",
		},
		{
			Pattern:  "+00(00)00000-0000",
			Expected: "+55(41)99433-6786",
		},
		{
			Pattern:  "00(00)000000000",
			Expected: "55(41)994336786",
		},
		{
			Pattern:  "",
			Expected: "5541994336786",
		},
	}

	for _, tc := range testCases {
		config.UserData.CellphonePattern = tc.Pattern
		assert.Equal(t, tc.Expected, FormatCellphone("5541994336786"))
	}
}
