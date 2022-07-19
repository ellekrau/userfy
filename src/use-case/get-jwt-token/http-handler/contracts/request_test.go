package contracts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateRequestContractWithValidContract(t *testing.T) {
	r := Request{
		Key: "Key",
	}

	assert.NoError(t, r.validateRequestContract())
}

func TestValidateRequestContractWithInvalidContract(t *testing.T) {
	testCase := "Empty key"
	expectedErrorMessage := "Key: 'Request.Key' Error:Field validation for 'Key' failed on the 'required' tag"
	request := Request{Key: ""}

	err := request.validateRequestContract()

	assert.Error(t, err)
	assert.Equal(t, expectedErrorMessage, err.Error(), testCase)
}
