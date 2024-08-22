package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerTodosRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoTodosRouter)
}

// 需认证的路由代码
//func registerTodosRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	api := apis.SysUser{}
//	r := v1.Group("/todo").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
//	{
//		r.GET("/", api.GetPage)
//	}
//}

// 不需要认证的路由接口
func registerNoTodosRouter(v1 *gin.RouterGroup) {
	api := apis.Todos{}
	todo := v1.Group("/todo")
	{
		todo.GET("/list", api.ListById)
		todo.POST("/add", api.AddTodo)
		todo.POST("/delete", api.DeleteToods)
		todo.GET("/find", api.GetById)
		todo.POST("/updata", api.UpdataTood)
	}
}
