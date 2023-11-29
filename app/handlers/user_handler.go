package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hyahm/golog"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/services"
	"github.com/rengensheng/backend/app/utils"
	"strings"
)

type UserHandler struct {
	userService *services.UserService
	roleService *services.RoleService
	menuService *services.MenuService
}

func NewUserHandler(userService *services.UserService, roleService *services.RoleService, menuService *services.MenuService) *UserHandler {
	return &UserHandler{userService: userService, roleService: roleService, menuService: menuService}
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	if val, has := c.Get("username"); has {
		user.CreatedBy = val.(string)
		user.UpdatedBy = val.(string)
	}
	userUp, err := handler.userService.CreateUser(&user)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(userUp, c)
}

func (handler *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := handler.userService.GetUserById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(user, c)
}

func (handler *UserHandler) UpdateUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	if val, has := c.Get("username"); has {
		user.UpdatedBy = val.(string)
	}
	userUp, err := handler.userService.UpdateUserById(id, &user)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(userUp, c)
}

func (handler *UserHandler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	err := handler.userService.DeleteUserById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(nil, c)
}

func (handler *UserHandler) UserLogin(c *gin.Context) {
	json := make(map[string]string)
	err := c.ShouldBindJSON(&json)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	if json["username"] == "" || json["password"] == "" {
		utils.ResultError(constant.FAILED, constant.USERNAME_PASSWORD_NULL, c)
		return
	}
	password := utils.MD5(json["password"])
	isExits := handler.userService.FindByUsernameAndPassword(json["username"], password)
	if !isExits {
		utils.ResultError(constant.FAILED, constant.USERNAME_PASSWORD_ERROR, c)
		return
	}
	user, err := handler.userService.GetUserByUsername(json["username"])
	var permission []string
	var roleValues []string
	roles, _ := handler.roleService.GetRoleByRoleValues(strings.Split(user.Role, ","))
	var menusIds []string
	for _, role := range roles {
		roleValues = append(roleValues, role.RoleValue)
		menus := strings.Split(role.Menu, ",")
		menusIds = append(menusIds, menus...)
	}
	menuList, err := handler.menuService.GetMenuListByIds(menusIds)
	if err == nil {
		for _, v := range menuList {
			if v.Permission != "" {
				permission = append(permission, v.Permission)
			}
		}
	}
	token, err := utils.GenerateToken(*user, permission, roleValues)
	c.SetCookie("token", token, 3600000, "/", "localhost", false, true)
	userLoginRes := models.UserLoginEntity{
		Desc:     user.Remark,
		RealName: user.Nickname,
		Token:    token,
		UserId:   user.Id,
		Username: user.Username,
	}
	utils.ResultSuccess(userLoginRes, c)
}

func (handler *UserHandler) GetCurrentUserInfo(c *gin.Context) {
	token := ""
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		golog.Info(err.Error())
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userClaims, err := utils.ParseToken(token)

	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	user, err := handler.userService.GetUserById(userClaims.ID)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(user, c)
}

func (handler *UserHandler) GetUserList(c *gin.Context) {
	var requestParams utils.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	menuList, err := handler.userService.GetUserList(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	count, err := handler.userService.GetUserListCount(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccessList(menuList, count, c)
}
