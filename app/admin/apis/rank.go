package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/models"
	"go-admin/app/admin/utils"
)

type simpleUser struct {
	UserId    int
	Username  string
	TimeTotal int
}

func Rank(c *gin.Context) {
	db := utils.GetDB()
	var users []models.SysUser
	db.Order("time_total desc").Find(&users)
	var simpleUsers []simpleUser
	for _, user := range users {
		simUser := simpleUser{
			UserId:    user.UserId,
			Username:  user.Username,
			TimeTotal: user.TimeTotal,
		}
		simpleUsers = append(simpleUsers, simUser)
	}

	c.JSON(200, simpleUsers)
}
