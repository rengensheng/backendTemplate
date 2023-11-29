package repositories

import (
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/utils"
	"xorm.io/xorm"
)

type UserRepository struct {
	db *xorm.Engine
}

func NewUserRepository(db *xorm.Engine) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if _, err := repo.db.Insert(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserById(id string) (*models.User, error) {
	user := &models.User{Id: id}
	if _, err := repo.db.Get(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) UpdateUserById(id string, user *models.User) (*models.User, error) {
	if _, err := repo.db.ID(id).Update(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) DeleteUserById(id string) error {
	user := &models.User{Id: id}
	if _, err := repo.db.Delete(&user); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserList(req *utils.Request) ([]*models.User, error) {
	var userList []*models.User
	req.DisposeRequest(repo.db.NewSession()).Find(&userList)
	if err := repo.db.Find(&userList); err != nil {
		return nil, err
	}
	return userList, nil
}

func (repo *UserRepository) GetUserListCount(req *utils.Request) (int64, error) {
	total, err := req.DisposeRequest(repo.db.NewSession()).Count(models.User{})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *UserRepository) GetUserAll() ([]*models.User, error) {
	var userList []*models.User
	if err := repo.db.Find(&userList); err != nil {
		return nil, err
	}
	return userList, nil
}

func (repo *UserRepository) GetUserAllCount() (int64, error) {
	total, err := repo.db.Count(&models.User{})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *UserRepository) GetUserExistByUsernameAndPassword(username, password string) bool {
	var user models.User
	_, err := repo.db.Where("account = ? and pwd = ?", username, password).Get(&user)
	if err != nil {
		return false
	}
	if user.Id == "" {
		return false
	}
	return true
}
