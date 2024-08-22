package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"time"
)

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/20 下午6:59
 **/

type ClockRoom struct {
	api.Api
}

/**
 * @加入自习室
 * @Param
 * @return
 * @Date 2024/8/20 下午7:02
 **/

func (e ClockRoom) EnterRoom(c *gin.Context) {
	s := service.ClockRoom{}
	req := models.ClockRoom{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "加入自习室成功！")
}

/**
 * @更新当前打卡情况
 * @Param
 * @return
 * @Date 2024/8/20 下午11:36
 **/

func (e ClockRoom) UpdataCur(c *gin.Context) {
	s := service.ClockRoom{}
	req := models.ClockRoom{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.UpdataCur(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "更新当前打卡情况成功！")
}

/**
 * @根据用户id查看自习室情况
 * @Param
 * @return
 * @Date 2024/8/21 上午12:06
 **/

func (e ClockRoom) GetByUserIdAndDate(c *gin.Context) {
	userid := c.Query("userid")
	date := c.Query("date")
	s := service.ClockRoom{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	data, err := s.GetByUserIdAndDate(userid, date)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查看自习室打卡情况成功！")
}

/**
 * @根据用户id查看自习记录
 * @Param
 * @return
 * @Date 2024/8/22 下午12:13
 **/

func (e ClockRoom) GetByUserId(c *gin.Context) {
	userid := c.Query("userid")
	s := service.ClockRoom{}

	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	data, err := s.GetByUserId(userid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查看自习室打卡情况成功！")
}

/**
 * @查询今日自习室（所有人）
 * @Param
 * @return
 * @Date 2024/8/20 下午11:53
 **/

func (e ClockRoom) ListRoom(c *gin.Context) {
	date := time.Now().Format("2006-01-02")
	s := service.ClockRoom{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	data, err := s.ListRoom(date)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查看今日打卡情况成功！")
}

/**
 * @根据ids删除自习记录
 * @Param
 * @return
 * @Date 2024/8/21 下午5:50
 **/

func (e *ClockRoom) DeleteRoom(c *gin.Context) {
	s := service.ClockRoom{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	idsStr := c.Query("ids") // 或 c.PostForm("ids")
	if idsStr == "" {
		e.Error(400, nil, "请传递待删除的 ids")
		return
	}

	err = s.DeleteRoom(idsStr)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "自习记录删除成功！")
}

/**
 * @根据id获取自习信息
 * @Param
 * @return
 * @Date 2024/8/21 下午5:57
 **/

func (e *ClockRoom) GetById(c *gin.Context) {
	s := service.ClockRoom{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	roomid := c.Query("roomid")
	data, err := s.GetById(roomid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查询打卡记录成功！")
}
