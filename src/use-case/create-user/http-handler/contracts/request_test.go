package contracts

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestCreateRequest(t *testing.T) {
	testCases := []struct {
		Name      string `json: "name"`
		Cellphone string `json: "cellphone"`
	}{
		{
			Name:      "name",
			Cellphone: "cellphone",
		},
	}

	for _, tc := range testCases {
		// Creates a gin test context
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

		// Converts the test request struct to json
		b, _ := json.Marshal(tc)

		// Injects the test request json in test context request body
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		r, err := CreateRequest(ctx)

		assert.Empty(t, r)
		assert.Error(t, err)
	}
}
