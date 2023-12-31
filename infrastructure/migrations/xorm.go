package migrations

import (
	"fmt"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/models"
	"github.com/rengensheng/backend/app/repositories"
	"github.com/rengensheng/backend/app/services"
	"github.com/rengensheng/backend/app/utils"
	"log"
	"xorm.io/xorm"
)

func Sync(engine *xorm.Engine) {
	log.Println(constant.START_SYNC_TABLE)
	err := engine.Sync(new(models.User))
	if err != nil {
		log.Println(constant.START_SYNC_DATA_ERROR, err.Error())
	}
	err = engine.Sync(new(models.Role))
	if err != nil {
		log.Println(constant.START_SYNC_DATA_ERROR, err.Error())
	}
	err = engine.Sync(new(models.Dept))
	if err != nil {
		log.Println(constant.START_SYNC_DATA_ERROR, err.Error())
	}
	err = engine.Sync(new(models.Menu))
	if err != nil {
		log.Println(constant.START_SYNC_DATA_ERROR, err.Error())
	}
	SyncTableData(engine)
}

func SyncTableData(engine *xorm.Engine) {
	log.Println("同步数据库结构.........")
	log.Println("添加默认管理员 admin----123456")
	userRepository := repositories.NewUserRepository(engine)
	// 查询用户表是否有数据，有数据不进行同步
	total, err := userRepository.GetUserAllCount()
	if err != nil {
		panic(constant.START_SYNC_DATA_ERROR + ":" + err.Error())
	}
	if total > 0 {
		log.Println("用户表已存在数据，不进行同步")
		return
	}
	roleRepository := repositories.NewRoleRepository(engine)
	deptRepository := repositories.NewDeptRepository(engine)
	menuRepository := repositories.NewMenuRepository(engine)
	userService := services.NewUserService(userRepository)
	roleService := services.NewRoleService(roleRepository)
	deptService := services.NewDeptService(deptRepository)
	menuService := services.NewMenuService(menuRepository)

	role := models.Role{
		RoleName:  "超级管理员",
		RoleValue: "admin",
		Status:    "0",
	}
	if err != nil {
		log.Println("添加默认管理员角色失败", err.Error())
	} else {
		log.Println("添加默认管理员角色成功...")
	}
	dept := models.Dept{
		DeptName: "默认部门",
		Status:   "0",
	}
	deptService.CreateDept(&dept)
	adminUser := models.User{
		Account:  "admin",
		Username: "超级管理员",
		Pwd:      utils.MD5("123456"),
		Nickname: "超级管理员",
		LoginId:  "admin",
		Role:     role.RoleValue,
		Dept:     dept.Id,
		Email:    "goylord2@gmail.com",
	}
	_, err = userService.CreateUser(&adminUser)
	if err != nil {
		log.Println("添加默认管理员账号失败", err.Error())
	} else {
		log.Println("添加默认管理员账号成功...")
	}
	dashboardRootMenu := models.Menu{
		MenuName:  "仪表盘",
		Status:    "0",
		Component: "LAYOUT",
		Icon:      "ant-design:dashboard-outlined",
		RoutePath: "/dashboard",
		Show:      "0",
		OrderNo:   1,
		Type:      "0",
		IsExt:     "0",
		Keepalive: "0",
	}
	analysisMenu := models.Menu{
		MenuName:   "分析页",
		Status:     "0",
		Component:  "dashboard/analysis/index.vue",
		Icon:       "ant-design:area-chart-outlined",
		RoutePath:  "analysis",
		ParentMenu: dashboardRootMenu.Id,
		Show:       "0",
		OrderNo:    0,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	workbenchMenu := models.Menu{
		MenuName:   "工作台",
		Status:     "0",
		Component:  "dashboard/workbench/index.vue",
		Icon:       "ant-design:calendar-twotone",
		RoutePath:  "workbench",
		ParentMenu: dashboardRootMenu.Id,
		Show:       "0",
		OrderNo:    1,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	// 创建菜单
	systemMenu := models.Menu{
		MenuName:  "系统管理",
		Status:    "0",
		Component: "LAYOUT",
		Icon:      "ant-design:setting-outlined",
		RoutePath: "/system",
		Show:      "0",
		OrderNo:   5,
		Type:      "0",
		IsExt:     "0",
		Keepalive: "0",
	}
	accountMenu := models.Menu{
		MenuName:   "账号管理",
		Status:     "0",
		Component:  "system/account/index.vue",
		Icon:       "ant-design:user-add-outlined",
		RoutePath:  "account",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    0,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	accountDetailMenu := models.Menu{
		MenuName:   "账号详情",
		Status:     "0",
		Component:  "system/account/AccountDetail.vue",
		Icon:       "ant-design:appstore-outlined",
		RoutePath:  "account_detail/:id",
		ParentMenu: systemMenu.Id,
		Show:       "1",
		OrderNo:    2,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	roleMenu := models.Menu{
		MenuName:   "角色管理",
		Status:     "0",
		Component:  "system/role/index.vue",
		Icon:       "ant-design:android-outlined",
		RoutePath:  "role",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    1,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	menuMenu := models.Menu{
		MenuName:   "菜单管理",
		Status:     "0",
		Component:  "system/menu/index.vue",
		Icon:       "ant-design:menu-outlined",
		RoutePath:  "menu",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    2,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	deptMenu := models.Menu{
		MenuName:   "部门管理",
		Status:     "0",
		Component:  "system/dept/index.vue",
		Icon:       "ant-design:deployment-unit-outlined",
		RoutePath:  "dept",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    3,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	menuService.CreateMenu(&dashboardRootMenu)
	menuService.CreateMenu(&analysisMenu)
	menuService.CreateMenu(&workbenchMenu)
	menuService.CreateMenu(&systemMenu)
	menuService.CreateMenu(&accountMenu)
	menuService.CreateMenu(&accountDetailMenu)
	menuService.CreateMenu(&roleMenu)
	menuService.CreateMenu(&menuMenu)
	menuService.CreateMenu(&deptMenu)
	role.Menu = fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s",
		dashboardRootMenu.Id, analysisMenu.Id, workbenchMenu.Id,
		systemMenu.Id, accountMenu.Id, roleMenu.Id, menuMenu.Id, deptMenu.Id)
	roleService.CreateRole(&role)
}
