package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/services"
	"github.com/rengensheng/backend/app/utils"
)

type MenuHandler struct {
	menuService *services.MenuService
}

func NewMenuHandler(menuService *services.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

func (handler *MenuHandler) CreateMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	if val, has := c.Get("username"); has {
		menu.CreatedBy = val.(string)
		menu.UpdatedBy = val.(string)
	}
	menuUp, err := handler.menuService.CreateMenu(&menu)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(menuUp, c)
}

func (handler *MenuHandler) GetMenuById(c *gin.Context) {
	id := c.Param("id")
	menu, err := handler.menuService.GetMenuById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(menu, c)
}

func (handler *MenuHandler) UpdateMenuById(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	if val, has := c.Get("username"); has {
		menu.UpdatedBy = val.(string)
	}
	menuUp, err := handler.menuService.UpdateMenuById(id, &menu)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(menuUp, c)
}

func (handler *MenuHandler) DeleteMenuById(c *gin.Context) {
	id := c.Param("id")
	err := handler.menuService.DeleteMenuById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(nil, c)
}

func (handler *MenuHandler) GetMenuList(c *gin.Context) {
	var requestParams utils.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	menuList, err := handler.menuService.GetMenuList(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	count, err := handler.menuService.GetMenuListCount(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccessList(menuList, count, c)
}
