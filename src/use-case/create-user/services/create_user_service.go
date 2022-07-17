package createuserservices

import (
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/domain"
	"github.com/ellekrau/mercafacil/repositories"
	createuserservicecontracts "github.com/ellekrau/mercafacil/use-case/create-user/services/contracts"
	customerror "github.com/ellekrau/mercafacil/utils/custom-error"
)

func CreateUser(input createuserservicecontracts.CreateUserServiceInput) (domain.User, *customerror.CustomError) {
	db := database.GetDatabase()

	userRepository := repositories.NewUserRepository(db)

	user := domain.NewUser(input.Name, input.Cellphone)
	if err := userRepository.CreateUser(&user); err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    "db_create_user",
			Message: err.Error(),
		}
	}

	return user, nil
}
