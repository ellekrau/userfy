package contracts

import (
	"fmt"
	customerror "github.com/ellekrau/userfy/utils/custom-error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	errInvalidRequestBodyCode    = "invalid_request_body"
	errInvalidRequestBodyMessage = "invalid request body: "
)

type Request struct {
	Key string `json:"name" validate:"required"`
}

func CreateRequest(c *gin.Context) (Request, *customerror.CustomError) {
	r := Request{Key: c.Query("key")}

	if err := r.validateRequestContract(); err != nil {
		return Request{}, &customerror.CustomError{
			Code:    errInvalidRequestBodyCode,
			Message: fmt.Sprint(errInvalidRequestBodyMessage, err.Error()),
		}
	}

	return r, nil
}

func (r Request) validateRequestContract() error {
	if err := validator.New().Struct(r); err != nil {
		return err
	}
	return nil
}
