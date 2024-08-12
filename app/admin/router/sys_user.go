package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/apis"
	"go-admin/app/admin/common"
	"go-admin/app/admin/utils"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoSysUserRouter)
}

// 需认证的路由代码
func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}
	r := v1.Group("/sys-user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("", api.Delete)
	}

	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
	{
		user.GET("/profile", api.GetProfile)
		user.POST("/avatar", api.InsetAvatar)
		user.PUT("/pwd/set", api.UpdatePwd)
		user.PUT("/pwd/reset", api.ResetPwd)
		user.PUT("/status", api.UpdateStatus)
	}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		v1auth.GET("/getinfo", api.GetInfo)
	}
}

// 不需要认证的路由接口
func registerNoSysUserRouter(v1 *gin.RouterGroup) {
	api := apis.SysUser{}
	user := v1.Group("/user")
	{
		user.POST("/getcode", utils.SendCode)
		user.POST("/register", api.Register)
		user.POST("/login", apis.Login)

		user.Use(common.AuthMiddleware())
		user.POST("/start", apis.Start)
		user.GET("/end", apis.End)

	}
}
