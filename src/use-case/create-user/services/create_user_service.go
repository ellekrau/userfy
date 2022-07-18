package createuserservices

import (
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/database/repositories/user-repository"
	"github.com/ellekrau/mercafacil/domain"
	createuserservicecontracts "github.com/ellekrau/mercafacil/use-case/create-user/services/contracts"
	customerror "github.com/ellekrau/mercafacil/utils/custom-error"
)

var errCodeDbCreateUser = "db_create_user"

func CreateUser(input createuserservicecontracts.CreateUserServiceInput) (domain.User, *customerror.CustomError) {
	db := database.GetDatabase()

	userRepository := userrepository.NewUserRepository(db)
	user := domain.NewUser(input.Name, input.Cellphone)

	if err := userRepository.CreateUser(&user); err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    errCodeDbCreateUser,
			Message: err.Error(),
		}
	}

	return user, nil
}
