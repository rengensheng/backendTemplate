package repositories

import (
	"github.com/hyahm/golog"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/utils"
	"xorm.io/xorm"
)

type RoleRepository struct {
	db *xorm.Engine
}

func NewRoleRepository(db *xorm.Engine) *RoleRepository {
	return &RoleRepository{db: db}
}

func (repo *RoleRepository) CreateRole(role *models.Role) (*models.Role, error) {
	if _, err := repo.db.Insert(&role); err != nil {
		return nil, err
	}
	return role, nil
}

func (repo *RoleRepository) GetRoleById(id string) (*models.Role, error) {
	role := &models.Role{Id: id}
	if _, err := repo.db.Get(role); err != nil {
		return nil, err
	}
	return role, nil
}

func (repo *RoleRepository) UpdateRoleById(id string, role *models.Role) (*models.Role, error) {
	if _, err := repo.db.ID(id).Update(&role); err != nil {
		return nil, err
	}
	return role, nil
}

func (repo *RoleRepository) DeleteRoleById(id string) error {
	role := &models.Role{Id: id}
	if _, err := repo.db.Delete(&role); err != nil {
		return err
	}
	return nil
}

func (repo *RoleRepository) GetRoleByRoleValues(ids []string) ([]*models.Role, error) {
	var roles []*models.Role
	err := repo.db.In("role_value", ids).Desc("id").Find(roles)
	if err != nil {
		golog.Info(err.Error())
		return nil, err
	}
	return roles, nil
}

func (repo *RoleRepository) GetRoleList(req *utils.Request) ([]*models.Role, error) {
	var roleList []*models.Role
	req.DisposeRequest(repo.db.NewSession()).Find(&roleList)
	if err := repo.db.Find(&roleList); err != nil {
		return nil, err
	}
	return roleList, nil
}

func (repo *RoleRepository) GetRoleListCount(req *utils.Request) (int64, error) {
	total, err := req.DisposeRequest(repo.db.NewSession()).Count(models.Role{})
	if err != nil {
		return 0, err
	}
	return total, nil
}
