package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/services"
	"github.com/rengensheng/backend/app/utils"
)

type DeptHandler struct {
	deptService *services.DeptService
}

func NewDeptHandler(deptService *services.DeptService) *DeptHandler {
	return &DeptHandler{deptService: deptService}
}

func (handler *DeptHandler) CreateDept(c *gin.Context) {
	var dept *models.Dept
	if err := c.ShouldBindJSON(dept); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	dept, err := handler.deptService.CreateDept(dept)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(dept, c)
}

func (handler *DeptHandler) GetDeptById(c *gin.Context) {
	id := c.Param("id")
	dept, err := handler.deptService.GetDeptById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(dept, c)
}

func (handler *DeptHandler) UpdateDeptById(c *gin.Context) {
	id := c.Param("id")
	var dept models.Dept
	if err := c.ShouldBindJSON(&dept); err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	deptUp, err := handler.deptService.UpdateDeptById(id, &dept)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(deptUp, c)
}

func (handler *DeptHandler) DeleteDeptById(c *gin.Context) {
	id := c.Param("id")
	err := handler.deptService.DeleteDeptById(id)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccess(nil, c)
}

func (handler *DeptHandler) GetDeptList(c *gin.Context) {
	var requestParams utils.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	deptList, err := handler.deptService.GetDeptList(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	count, err := handler.deptService.GetDeptListCount(&requestParams)
	if err != nil {
		utils.ResultError(constant.FAILED, err.Error(), c)
		return
	}
	utils.ResultSuccessList(deptList, count, c)
}
