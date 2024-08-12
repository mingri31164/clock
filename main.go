package main

import (
	"go-admin/app/admin/common"
	"go-admin/app/admin/utils"
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 521386980 进入技术交流群 请先star，谢谢！
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	db := utils.InitDB()
	//获取底层数据库连接
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	defer sqlDB.Close()
	common.StartScheduledDeletion(db)

	cmd.Execute()
}
