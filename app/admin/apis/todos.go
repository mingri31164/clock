package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/admin/common"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/app/admin/utils"
	"net/http"
	"strconv"
)

type Todos struct {
	api.Api
}

// 初始化redis配置
var redisdb = utils.InitRedis()

/**
 * @根据用户id查询待办
 * @Param
 * @return
 * @Date 2024/8/19 16:33
 **/

func (e Todos) ListById(c *gin.Context) {
	userid := c.Query("userId")
	var todolist []*models.Todos
	//构造redis中的key：todos_userid
	todokey := "todos_" + userid
	//查询redis中是否存在该用户的待办数据
	todoData, err := redisdb.Get(todokey).Result()
	//如果存在，直接返回，无需查询数据库
	if todoData != "" {
		err = models.DecodeTodos([]byte(todoData), &todolist) //反序列化为[]*models.Todos
		if err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		common.ResOK(c, "查询待办成功！", todolist)
		return
	}
	//如果不存在，查询数据库，将查询结果存入redis
	s := service.Todos{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	data, err := s.ListById(userid)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//将data转换为字符串存储到redis中
	redisTodos, err := json.Marshal(data)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	cmd := redisdb.Set(todokey, redisTodos, 0)
	if err := cmd.Err(); err != nil {
		common.ResErr(c, err.Error())
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
	//清除redis缓存
	key := "todos_" + strconv.Itoa(req.UserID)
	redisdb.Del(key)
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
	userid := c.Query("userId")
	idsStr := c.Query("todoIds") // 或 c.PostForm("ids")
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
	//清除redis缓存
	key := "todos_" + userid
	redisdb.Del(key)

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

	idStr := c.Query("todoId")
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
	data, err := s.GetById(req.TodoId)
	if err != nil {
		e.Logger.Error(err)
		e.Error(400, err, err.Error())
		return
	}
	//清除redis缓存
	key := "todos_" + strconv.Itoa(data.UserID)
	redisdb.Del(key)
	e.OK(nil, "修改成功！")
}
