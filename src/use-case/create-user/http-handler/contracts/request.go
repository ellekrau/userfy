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
	Name      string `json:"name" validate:"required"`
	Cellphone string `json:"cellphone" validate:"required,number,len=13"`
}

func CreateRequest(c *gin.Context) (Request, *customerror.CustomError) {
	r := Request{}
	if err := c.ShouldBindJSON(&r); err != nil {
		return Request{}, &customerror.CustomError{
			Code:    errInvalidRequestBodyCode,
			Message: fmt.Sprint(errInvalidRequestBodyMessage, err.Error()),
		}
	}

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
