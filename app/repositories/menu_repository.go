package repositories

import (
	"github.com/hyahm/golog"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/utils"
	"xorm.io/xorm"
)

type MenuRepository struct {
	db *xorm.Engine
}

func NewMenuRepository(db *xorm.Engine) *MenuRepository {
	return &MenuRepository{db: db}
}

func (repo *MenuRepository) CreateMenu(menu *models.Menu) (*models.Menu, error) {
	if _, err := repo.db.Insert(menu); err != nil {
		return nil, err
	}
	return menu, nil
}

func (repo *MenuRepository) GetMenuById(id string) (*models.Menu, error) {
	menu := &models.Menu{Id: id}
	if _, err := repo.db.Get(menu); err != nil {
		return nil, err
	}
	return menu, nil
}

func (repo *MenuRepository) UpdateMenuById(id string, menu *models.Menu) (*models.Menu, error) {
	if _, err := repo.db.ID(id).Update(menu); err != nil {
		return nil, err
	}
	return menu, nil
}

func (repo *MenuRepository) DeleteMenuById(id string) error {
	menu := &models.Menu{Id: id}
	if _, err := repo.db.Delete(menu); err != nil {
		return err
	}
	return nil
}

func (repo *MenuRepository) GetMenuListByIds(ids []string) ([]models.Menu, error) {
	var menus []models.Menu
	err := repo.db.In("id", ids).Asc("order_no").Find(&menus)
	if err != nil {
		golog.Info(err.Error())
		return nil, err
	}
	return menus, nil
}

func (repo *MenuRepository) GetMenuList(req *utils.Request) ([]models.Menu, error) {
	var menuList []models.Menu
	err := req.DisposeRequest(repo.db.NewSession()).Find(&menuList)
	if err != nil {
		return nil, err
	}
	return menuList, nil
}

func (repo *MenuRepository) GetMenuListCount(req *utils.Request) (int64, error) {
	total, err := req.DisposeRequest(repo.db.NewSession()).Count(models.Menu{})
	if err != nil {
		return 0, err
	}
	return total, nil
}
