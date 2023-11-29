# 基于gin的简易后台开发模板

## 项目结构
```
| |____app
| | |____app.go
| | |____config
| | | |____config.go
| | |____constant
| | | |____error_msg.go
| | | |____http_status.go
| | |____handlers
| | | |____dept_handler.go
| | | |____menu_handler.go
| | | |____role_handler.go
| | | |____user_handler.go
| | |____middlewares
| | | |____auth.go
| | |____models
| | | |____dept.go
| | | |____menu.go
| | | |____role.go
| | | |____user.go
| | |____repositories
| | | |____dept_repository.go
| | | |____menu_repository.go
| | | |____role_repository.go
| | | |____user_repository.go
| | |____routes
| | | |____dept_route.go
| | | |____setup.go
| | | |____user_route.go
| | |____services
| | | |____dept_service.go
| | | |____menu_service.go
| | | |____role_service.go
| | | |____user_service.go
| | |____utils
| | | |____request.go
| | | |____result.go
| | | |____user.go
| | | |____utils.go
| |____backend
| |____codegen
| | |____generator.go
| | |____template
| |____config.yaml
| |____go.mod
| |____go.sum
| |____infrastructure
| | |____database
| | | |____database.go
| | |____migrations
| | | |____xorm.go
| |____logger
| | |____logger.go
| |____main.go
| |____README.md
| |____scripts
| |____tests

```