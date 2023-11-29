package services

import (
	"github.com/golang-module/carbon"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	user.Id = utils.GetUUID()
	user.CreatedTime = carbon.Now().ToDateTimeString()
	user.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.userRepository.CreateUser(user)
}

func (service *UserService) GetUserById(id string) (*models.User, error) {
	return service.userRepository.GetUserById(id)
}

func (service *UserService) UpdateUserById(id string, user *models.User) (*models.User, error) {
	user.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.userRepository.UpdateUserById(id, user)
}

func (service *UserService) DeleteUserById(id string) error {
	return service.userRepository.DeleteUserById(id)
}

func (service *UserService) GetUserList(request *utils.Request) ([]*models.User, error) {
	return service.userRepository.GetUserList(request)
}

func (service *UserService) GetUserListCount(request *utils.Request) (int64, error) {
	return service.userRepository.GetUserListCount(request)
}

func (service *UserService) FindByUsernameAndPassword(username string, password string) bool {
	return service.userRepository.GetUserExistByUsernameAndPassword(username, password)
}
