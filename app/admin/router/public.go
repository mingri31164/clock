package router

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/21 下午8:09
 **/
import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/utils"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerClockRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoPublicRouter)
}

// 需认证的路由代码
//func registerClockRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	api := apis.SysUser{}
//	r := v1.Group("/clock").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole()).Use(actions.PermissionAction())
//	{
//		r.GET("/", api.GetPage)
//	}
//}

// 不需要认证的路由接口
func registerNoPublicRouter(v1 *gin.RouterGroup) {
	public := v1.Group("/public")
	{
		public.GET("/getcode", utils.SendCode)
	}

}
