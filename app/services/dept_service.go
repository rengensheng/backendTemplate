package services

import (
	"github.com/golang-module/carbon"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/utils"
)

type DeptService struct {
	deptRepository *repositories.DeptRepository
}

func NewDeptService(deptRepository *repositories.DeptRepository) *DeptService {
	return &DeptService{deptRepository: deptRepository}
}

func (service *DeptService) CreateDept(dept *models.Dept) (*models.Dept, error) {
	dept.Id = utils.GetUUID()
	dept.CreatedTime = carbon.Now().ToDateTimeString()
	dept.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.deptRepository.CreateDept(dept)
}

func (service *DeptService) GetDeptById(id string) (*models.Dept, error) {
	return service.deptRepository.GetDeptById(id)
}

func (service *DeptService) UpdateDeptById(id string, dept *models.Dept) (*models.Dept, error) {
	dept.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.deptRepository.UpdateDeptById(id, dept)
}

func (service *DeptService) DeleteDeptById(id string) error {
	return service.deptRepository.DeleteDeptById(id)
}

func (service *DeptService) GetDeptList(request *utils.Request) ([]models.Dept, error) {
	return service.deptRepository.GetDeptList(request)
}

func (service *DeptService) GetDeptListCount(request *utils.Request) (int64, error) {
	return service.deptRepository.GetDeptListCount(request)
}
