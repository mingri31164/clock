package router

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午9:27
 **/

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerClockRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerNoClockRouter)
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
func registerNoClockRouter(v1 *gin.RouterGroup) {
	api := apis.Clock{}
	roomApi := apis.ClockRoom{}
	clock := v1.Group("/clock")
	{
		clock.POST("/start", api.AddClock)
		clock.GET("/listByDate", api.ListByDate)
		clock.POST("/end", api.EndClock)
		clock.GET("/listByUserId", api.ListByUserId)
		clock.GET("/getById", api.GetById)
		clock.POST("/delete", api.DeleteClock)
		clock.POST("/updata", api.UpdataClock)
	}

	room := v1.Group("/room")
	{
		room.POST("/add", roomApi.EnterRoom)
		room.POST("/delete", roomApi.DeleteRoom)
		room.POST("/updata", roomApi.UpdataCur)
		room.GET("/getById", roomApi.GetById)
		room.GET("/getByUserId", roomApi.GetByUserId)
		room.GET("/getByUserAndDate", roomApi.GetByUserIdAndDate)
		room.GET("/list", roomApi.ListRoom)
		room.GET("/listFinishTodes", roomApi.ListFinishTodes)
	}

}
