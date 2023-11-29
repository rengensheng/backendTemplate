package repositories

import (
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/utils"
	"xorm.io/xorm"
)

type DeptRepository struct {
	db *xorm.Engine
}

func NewDeptRepository(db *xorm.Engine) *DeptRepository {
	return &DeptRepository{db: db}
}

func (repo *DeptRepository) CreateDept(dept *models.Dept) (*models.Dept, error) {
	if _, err := repo.db.Insert(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (repo *DeptRepository) GetDeptById(id string) (*models.Dept, error) {
	dept := &models.Dept{Id: id}
	if _, err := repo.db.Get(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (repo *DeptRepository) UpdateDeptById(id string, dept *models.Dept) (*models.Dept, error) {
	if _, err := repo.db.ID(id).Update(dept); err != nil {
		return nil, err
	}
	return dept, nil
}

func (repo *DeptRepository) DeleteDeptById(id string) error {
	dept := &models.Dept{Id: id}
	if _, err := repo.db.Delete(dept); err != nil {
		return err
	}
	return nil
}

func (repo *DeptRepository) GetDeptList(req *utils.Request) ([]models.Dept, error) {
	var deptList []models.Dept
	err := req.DisposeRequest(repo.db.NewSession()).Find(&deptList)
	if err != nil {
		return nil, err
	}
	return deptList, nil
}

func (repo *DeptRepository) GetDeptListCount(req *utils.Request) (int64, error) {
	total, err := req.DisposeRequest(repo.db.NewSession()).Count(models.Dept{})
	if err != nil {
		return 0, err
	}
	return total, nil
}
