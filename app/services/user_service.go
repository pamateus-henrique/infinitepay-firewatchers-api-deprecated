package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/models"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/app/queries"
	"github.com/pamateus-henrique/infinitepay-firewatchers-api/pkg/utils"
)



type UserService struct {
	queries *queries.UserQueries
	validator *validator.Validate
}


func CreateUserService(queries *queries.UserQueries, ) *UserService{
	return &UserService{
		queries: queries,
		validator: validator.New(),
	}
}


func (us *UserService) CreateUser(userModel *models.User) error {

	if err := us.queries.CreateUser(userModel); err != nil {
		return err
	}

	return nil 
}


func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	
	user, err := us.queries.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil

}


func (us *UserService) CreateUserModel(registerModel *models.Register) *models.User {
	
	userModel := models.User{}

	userModel.Name = registerModel.Name
	userModel.Email = registerModel.Email
	userModel.Password = utils.GeneratePassword(registerModel.Password)
	userModel.Avatar_url = "abc"
	userModel.Role = "Viewer"
	userModel.Team = "None"

	return &userModel
}

func (us *UserService) ValidateRegisterModel(authModel *models.Register) error {
	validate := validator.New()

	return validate.Struct(authModel)

}

func (us *UserService) ValidateLoginModel(authModel *models.Login) error {
	validate := validator.New()

	return validate.Struct(authModel)

}


func (us *UserService) ValidateUserModel(userModel *models.User) error {
	validate := validator.New()

	return validate.Struct(userModel)

}

