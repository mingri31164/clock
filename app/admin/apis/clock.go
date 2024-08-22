package apis

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午6:23
 **/

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"time"
)

type Clock struct {
	api.Api
}

/**
 * @开始打卡
 * @Param
 * @return
 * @Date 2024/8/19 下午9:41
 **/

func (e Clock) AddClock(c *gin.Context) {
	s := service.Clock{}
	t := service.Todos{}
	clock := service.Clock{}
	u := service.SysUser{}
	r := service.ClockRoom{}

	req := dto.AddClock{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		MakeService(&t.Service).
		MakeService(&clock.Service).
		MakeService(&u.Service).
		MakeService(&r.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Insert(&req, &t, &clock, &u, &r)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "打卡成功")
}

/**
 * @删除打卡记录
 * @Param
 * @return
 * @Date 2024/8/19 下午9:41
 **/

func (e *Clock) DeleteClock(c *gin.Context) {
	s := service.Clock{}
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

	err = s.Delete(idsStr)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "删除打卡记录成功！")
}

/**
 * @结束本次打卡
 * @Param
 * @return
 * @Date 2024/8/19 下午10:29
 **/

func (e Clock) EndClock(c *gin.Context) {
	clockid := c.Query("clockid")
	s := service.Clock{}
	u := service.SysUser{}
	t := service.Todos{}
	r := service.ClockRoom{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeService(&u.Service).
		MakeService(&t.Service).
		MakeService(&r.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.EndClock(clockid, &u, &t, &r)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "结束打卡成功！")
}

/**
 * @根据日期查询打卡记录
 * @Param
 * @return
 * @Date 2024/8/19 下午9:46
 **/

func (e *Clock) ListByDate(c *gin.Context) {
	s := service.Clock{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	clockDateString := c.Query("clockDate")
	// 将clockDateString转换为time.Time类型
	clockDate, err := time.Parse("2006-01-02", clockDateString)
	if err != nil {
		e.Logger.Errorf("error parsing clockDate: %s", err)
		// 在这里处理错误，例如发送错误响应给客户端
		return
	}

	data, err := s.ListByDate(clockDate)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查询记录成功！")
}

/**
 * @根据用户id查询所有打卡记录
 * @Param
 * @return
 * @Date 2024/8/20 下午1:42
 **/

func (e *Clock) ListByUserId(c *gin.Context) {
	s := service.Clock{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	userid := c.Query("userid")

	data, err := s.ListByUserId(userid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查询用户打卡记录成功！")
}

/**
 * @根据id获取记录信息
 * @Param
 * @return
 * @Date 2024/8/20 下午5:54
 **/

func (e *Clock) GetById(c *gin.Context) {
	s := service.Clock{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	userid := c.Query("clockid")
	data, err := s.GetById(userid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查询打卡记录成功！")
}

/**
 * @更新打卡记录
 * @Param
 * @return
 * @Date 2024/8/21 下午2:08
 **/

func (e Clock) UpdataClock(c *gin.Context) {
	s := service.Clock{}
	req := models.Clock{}
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

	err = s.UpdataClock(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "更新打卡记录成功！")
}
