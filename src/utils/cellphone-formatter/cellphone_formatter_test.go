package cellphoneformatter

import (
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatCellphone(t *testing.T) {
	testCases := []struct {
		Pattern  string
		Expected string
	}{
		{
			Pattern:  "+XX (XX) XXXXX-XXXX",
			Expected: "+55 (41) 99433-0786",
		},
		{
			Pattern:  "+XX(XX)XXXXX-XXXX",
			Expected: "+55(41)99433-0786",
		},
		{
			Pattern:  "XX(XX)XXXXXXXXX",
			Expected: "55(41)994330786",
		},
		{
			Pattern:  "",
			Expected: "5541994330786",
		},
	}

	for _, tc := range testCases {
		userConfig := clientconfig.User{
			Cellphone: tc.Pattern,
		}

		assert.Equal(t, tc.Expected, FormatCellphone(userConfig, "5541994330786"))
	}
}
