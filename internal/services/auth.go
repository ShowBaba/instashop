package services

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"instashop/internal/common"
	"instashop/internal/dtos"
	"instashop/internal/repositories"
	"instashop/internal/utils"
	"instashop/models"
)

type AuthClient interface {
	Login(input dtos.LoginDTO) (*dtos.LoginResp, *common.RestErr)
	Signup(input dtos.SignUpDTO) (*models.GetUser, *common.RestErr)
}

type AuthService struct {
	userRepo *repositories.UserRepository
	restErr  *common.RestErr
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	restErr *common.RestErr,
) AuthClient {
	return &AuthService{
		userRepo,
		restErr,
	}
}

func (a *AuthService) Login(input dtos.LoginDTO) (*dtos.LoginResp, *common.RestErr) {
	user, exist, err := a.userRepo.FetchOne(models.User{Email: input.Email})
	if err != nil {
		log.Error(zap.Error(err))
		return nil, a.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	if !exist {
		return nil, a.restErr.BadRequest(common.ErrUserWithEmailNotFound)
	}

	passwordMatch, err := utils.PasswordMatches(input.Password, user.Password)
	if err != nil {
		log.Error(zap.Error(err))
		return nil, a.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	if !passwordMatch {
		return nil, a.restErr.BadRequest(common.ErrInvalidPassword)
	}

	token, err := utils.GenerateToken(utils.GetConfig().JWTSecretKey, input.Email, user.ID, func() string {
		if *user.IsAdmin {
			return "admin"
		}
		return "user"
	}())
	if err != nil {
		log.Error(zap.Error(err))
		return nil, a.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	return &dtos.LoginResp{Token: token}, nil
}

func (a *AuthService) Signup(input dtos.SignUpDTO) (*models.GetUser, *common.RestErr) {
	_, exist, err := a.userRepo.FetchOne(models.User{Email: input.Email})
	if err != nil {
		log.Error(zap.Error(err))
		return nil, a.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	if exist {
		return nil, a.restErr.BadRequest(common.ErrEmailAlreadyInUse)
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Error(zap.Error(err))
		return nil, a.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	newUser := models.User{
		Email:       input.Email,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Password:    hashedPassword,
		PhoneNumber: input.PhoneNumber,
	}

	err = a.userRepo.Create(&newUser)
	if err != nil {
		log.Error(zap.Error(err))
		return nil, a.restErr.ServerError(common.ErrSomethingWentWrong)
	}

	// TODO: send verification email

	return newUser.ToGetUser(), nil
}
