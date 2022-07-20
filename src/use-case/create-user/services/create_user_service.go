package createuserservices

import (
	"github.com/ellekrau/mercafacil/config"
	"github.com/ellekrau/mercafacil/database"
	"github.com/ellekrau/mercafacil/database/repositories/user-repository"
	"github.com/ellekrau/mercafacil/domain"
	createuserservicecontracts "github.com/ellekrau/mercafacil/use-case/create-user/services/contracts"
	customerror "github.com/ellekrau/mercafacil/utils/custom-error"
)

var errCodeDbCreateUser = "db_create_user"

func CreateUser(input createuserservicecontracts.CreateUserServiceInput) (domain.User, *customerror.CustomError) {
	client, err := config.GetClient(input.ClientKey)
	if err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    "client_not_configured",
			Message: "Client not configured.",
		}
	}

	db, err := database.GetDatabaseByClientKey(input.ClientKey)
	if err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    "get_database",
			Message: err.Error(),
		}
	}

	userRepository := userrepository.NewUserRepository(client, db)
	user := domain.NewUser(client.DataPattern.User, input.Name, input.Cellphone)

	if err = userRepository.CreateUser(&user); err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    errCodeDbCreateUser,
			Message: err.Error(),
		}
	}

	return user, nil
}
