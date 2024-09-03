package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	//routerCheckRole = append(routerCheckRole, registerClockRouter)
	routerNoCheckRole = append(routerNoCheckRole, rankRouter)
}

func rankRouter(v1 *gin.RouterGroup) {
	rank := v1.Group("/rank")
	{
		rank.GET("/listAll", apis.Rank)
	}
}
