package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"net/http"
	"strconv"
)

type Todos struct {
	api.Api
}

/**
 * @根据用户id查询待办
 * @Param
 * @return
 * @Date 2024/8/19 16:33
 **/

func (e Todos) ListById(c *gin.Context) {
	userid := c.Query("userid")
	s := service.Todos{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	data, err := s.ListById(userid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	if len(data) == 0 {
		e.Error(400, nil, "待办不存在！")
		return
	}
	e.OK(data, "查询待办成功！")
}

/**
 * @新增待办
 * @Param
 * @return
 * @Date 2024/8/19 16:34
 **/

func (e Todos) AddTodo(c *gin.Context) {
	s := service.Todos{}
	req := dto.AddTodo{}
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
	e.OK(nil, "创建成功")
}

/**
 * @根据ids删除待办
 * @Param
 * @return
 * @Date 2024/8/19 16:34
 **/

func (e *Todos) DeleteToods(c *gin.Context) {
	s := service.Todos{}
	r := service.ClockRoom{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		MakeService(&r.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	userid := c.Query("userid")
	idsStr := c.Query("ids") // 或 c.PostForm("ids")
	if idsStr == "" {
		e.Error(400, nil, "请传递待删除的 ids")
		return
	}

	err = s.Delete(userid, idsStr, &r)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "待办删除成功！")
}

/**
 * @根据id查询待办(数据回显)
 * @Param
 * @return
 * @Date 2024/8/19 16:37
 **/

func (e Todos) GetById(c *gin.Context) {
	s := service.Todos{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	idStr := c.Query("todoid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// 处理错误，例如发送错误响应给客户端
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	data, err := s.GetById(id)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(data, "查询待办成功！")
}

/**
 * @修改待办
 * @Param
 * @return
 * @Date 2024/8/19 16:35
 **/

func (e Todos) UpdataTood(c *gin.Context) {
	s := service.Todos{}
	req := models.Todos{}
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

	err = s.UpdataTodo(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, "修改成功！")
}
