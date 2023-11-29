package services

import (
	"github.com/golang-module/carbon"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/utils"
)

type MenuService struct {
	menuRepository *repositories.MenuRepository
}

func NewMenuService(menuRepository *repositories.MenuRepository) *MenuService {
	return &MenuService{menuRepository: menuRepository}
}

func (service *MenuService) CreateMenu(menu *models.Menu) (*models.Menu, error) {
	menu.Id = utils.GetUUID()
	menu.CreatedTime = carbon.Now().ToDateTimeString()
	menu.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.menuRepository.CreateMenu(menu)
}

func (service *MenuService) GetMenuById(id string) (*models.Menu, error) {
	return service.menuRepository.GetMenuById(id)
}

func (service *MenuService) UpdateMenuById(id string, menu *models.Menu) (*models.Menu, error) {
	menu.UpdatedTime = carbon.Now().ToDateTimeString()
	return service.menuRepository.UpdateMenuById(id, menu)
}

func (service *MenuService) DeleteMenuById(id string) error {
	return service.menuRepository.DeleteMenuById(id)
}

func (service *MenuService) GetMenuListByIds(ids []string) ([]models.Menu, error) {
	return service.menuRepository.GetMenuListByIds(ids)
}

func (service *MenuService) GetMenuList(request *utils.Request) ([]models.Menu, error) {
	return service.menuRepository.GetMenuList(request)
}

func (service *MenuService) GetMenuListCount(request *utils.Request) (int64, error) {
	return service.menuRepository.GetMenuListCount(request)
}
