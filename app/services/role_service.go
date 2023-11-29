package services

import (
	"github.com/golang-module/carbon"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/utils"
)

type RoleService struct {
	roleRepository *repositories.RoleRepository
}

func NewRoleService(roleRepository *repositories.RoleRepository) *RoleService {
	return &RoleService{roleRepository: roleRepository}
}

func (service *RoleService) CreateRole(role *models.Role) (*models.Role, error) {
	role.Id = utils.GetUUID()
	role.CreatedTime = carbon.Now().ToDateTimeString()
	role.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.roleRepository.CreateRole(role)
}

func (service *RoleService) GetRoleById(id string) (*models.Role, error) {
	return service.roleRepository.GetRoleById(id)
}

func (service *RoleService) UpdateRoleById(id string, role *models.Role) (*models.Role, error) {
	role.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.roleRepository.UpdateRoleById(id, role)
}

func (service *RoleService) DeleteRoleById(id string) error {
	return service.roleRepository.DeleteRoleById(id)
}

func (service *RoleService) GetRoleByRoleValues(roleValues []string) ([]*models.Role, error) {
	return service.roleRepository.GetRoleByRoleValues(roleValues)
}

func (service *RoleService) GetRoleList(request *utils.Request) ([]*models.Role, error) {
	return service.roleRepository.GetRoleList(request)
}

func (service *RoleService) GetRoleListCount(request *utils.Request) (int64, error) {
	return service.roleRepository.GetRoleListCount(request)
}
