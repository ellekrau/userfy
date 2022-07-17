package contracts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRequestContract(t *testing.T) {
	testCases := []struct {
		TestCase     string
		Request      Request
		ErrorMessage string
	}{
		{
			TestCase: "Cellphone with letters",
			Request: Request{
				Name:      "name",
				Cellphone: "cellphone",
			},
			ErrorMessage: "Key: 'Request.Cellphone' Error:Field validation for 'Cellphone' failed on the 'number' tag",
		},
		{
			TestCase: "Cellphone with letters and numbers",
			Request: Request{
				Name:      "name",
				Cellphone: "479928342a",
			},
			ErrorMessage: "Key: 'Request.Cellphone' Error:Field validation for 'Cellphone' failed on the 'number' tag",
		},
		{
			TestCase: "Empty name",
			Request: Request{
				Name:      "",
				Cellphone: "999999999",
			},
			ErrorMessage: "Key: 'Request.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
		{
			TestCase: "Empty cellphone",
			Request: Request{
				Name:      "Name",
				Cellphone: "",
			},
			ErrorMessage: "Key: 'Request.Cellphone' Error:Field validation for 'Cellphone' failed on the 'required' tag",
		},
		{
			TestCase: "Cellphone with insufficient digits (10)",
			Request: Request{
				Name:      "Name",
				Cellphone: "1234567890",
			},
			ErrorMessage: "Key: 'Request.Cellphone' Error:Field validation for 'Cellphone' failed on the 'len' tag",
		},
	}

	for _, tc := range testCases {
		err := tc.Request.validateRequestContract()
		assert.Equal(t, tc.ErrorMessage, err.Error(), tc.TestCase)
	}
}
