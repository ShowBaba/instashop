package services

import (
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"instashop/internal/common"
	"instashop/internal/repositories"
	"instashop/models"
)

type UserClient interface {
	GetUserDetails(userId uint) (*models.GetUser, *common.RestErr)
}

type UserService struct {
	userRepo *repositories.UserRepository
	restErr  *common.RestErr
}

func NewUserService(
	userRepo *repositories.UserRepository,
	restErr *common.RestErr,
) UserClient {
	return &UserService{
		userRepo,
		restErr,
	}
}

func (u *UserService) GetUserDetails(userId uint) (*models.GetUser, *common.RestErr) {
	user, exist, err := u.userRepo.FetchOne(models.User{ID: userId})
	if err != nil {
		log.Error(zap.Error(err))
		return nil, u.restErr.ServerError(common.ErrSomethingWentWrong)
	}
	if !exist {
		return nil, u.restErr.ServerError(common.ErrUserWithEmailNotFound)
	}

	return user.ToGetUser(), nil
}
