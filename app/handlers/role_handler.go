package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/services"
	"github.com/rengensheng/backend/app/utils"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (handler *RoleHandler) CreateRole(c *gin.Context) {
	var role *models.Role
	if err := c.ShouldBindJSON(role); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	role, err := handler.roleService.CreateRole(role)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(role, c)
}

func (handler *RoleHandler) GetRoleById(c *gin.Context) {
	id := c.Param("id")
	role, err := handler.roleService.GetRoleById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(role, c)
}

func (handler *RoleHandler) UpdateRoleById(c *gin.Context) {
	id := c.Param("id")
	var role *models.Role
	if err := c.ShouldBindJSON(role); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	role, err := handler.roleService.UpdateRoleById(id, role)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(role, c)
}

func (handler *RoleHandler) DeleteRoleById(c *gin.Context) {
	id := c.Param("id")
	err := handler.roleService.DeleteRoleById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(nil, c)
}

func (handler *RoleHandler) GetRoleList(c *gin.Context) {
	var requestParams utils.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	menuList, err := handler.roleService.GetRoleList(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	count, err := handler.roleService.GetRoleListCount(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccessList(menuList, count, c)
}
