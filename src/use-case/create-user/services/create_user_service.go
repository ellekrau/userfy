package createuserservices

import (
	clientconfig "github.com/ellekrau/userfy/config/client-config"
	"github.com/ellekrau/userfy/database"
	userrepository "github.com/ellekrau/userfy/database/repositories/user-repository"
	"github.com/ellekrau/userfy/domain"
	createuserservicecontracts "github.com/ellekrau/userfy/use-case/create-user/services/contracts"
	customerror "github.com/ellekrau/userfy/utils/custom-error"
)

var (
	errCodeDbCreateUser        = "db_create_user"
	errCodeDbGetDatabase       = "get_database"
	errCodeDbGetUserRepository = "get_user_repository"
	errClientNotConfigured     = "client_not_configured"
)

func CreateUser(input createuserservicecontracts.CreateUserServiceInput) (domain.User, *customerror.CustomError) {
	client, err := clientconfig.GetClient(input.ClientKey)
	if err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    errClientNotConfigured,
			Message: err.Error(),
		}
	}

	db, err := database.GetDatabaseByClientKey(input.ClientKey)
	if err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    errCodeDbGetDatabase,
			Message: err.Error(),
		}
	}

	userRepository, err := userrepository.NewUserRepository(client, db)
	if err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    errCodeDbGetUserRepository,
			Message: err.Error(),
		}
	}

	user := domain.NewUser(client.DataPattern.User, input.Name, input.Cellphone)
	if err = userRepository.CreateUser(&user); err != nil {
		return domain.User{}, &customerror.CustomError{
			Code:    errCodeDbCreateUser,
			Message: err.Error(),
		}
	}

	return user, nil
}
