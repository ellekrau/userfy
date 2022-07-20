package createuserservicecontracts

type CreateUserServiceInput struct {
	ClientKey string
	Name      string
	Cellphone string
}

func NewCreateUserServiceInput(clientKey, name, cellphone string) CreateUserServiceInput {
	return CreateUserServiceInput{
		ClientKey: clientKey,
		Name:      name,
		Cellphone: cellphone,
	}
}
