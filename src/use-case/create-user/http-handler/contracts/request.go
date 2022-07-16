package contracts

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Name      string `json:"name" validate:"required"`
	Cellphone string `json:"cellphone" validate:"required,number"`
}

func CreateRequest(c *gin.Context) (Request, error) {
	r := Request{}
	if err := c.ShouldBindJSON(&r); err != nil {
		return Request{}, err
	}

	if err := r.validate(); err != nil {
		return Request{}, err
	}

	return r, nil
}

func (r Request) validate() error {
	if err := validator.New().Struct(r); err != nil {
		return err
	}
	return nil
}
