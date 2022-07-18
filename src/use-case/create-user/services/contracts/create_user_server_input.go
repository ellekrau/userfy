package createuserservicecontracts

type CreateUserServiceInput struct {
	Name      string
	Cellphone string
}

func NewCreateUserServiceInput(name, cellphone string) CreateUserServiceInput {
	return CreateUserServiceInput{
		Name:      name,
		Cellphone: cellphone,
	}
}
