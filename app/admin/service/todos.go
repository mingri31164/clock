package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"strconv"
	"strings"
	"time"
)

type Todos struct {
	service.Service
}

/**
 * @新增待办
 * @Param
 * @return
 * @Date 2024/8/19 下午10:22
 **/

func (e *Todos) Insert(c *dto.AddTodo) error {
	var err error
	var data models.Todos
	var user models.SysUser
	var i int64

	err = e.Orm.Model(&user).Where("user_id = ?", c.UserID).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i == 0 {
		err := errors.New("该用户不存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	err = e.Orm.Model(&data).Where("name = ?", c.Name).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("该待办已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

/**
 * @根据用户id查询待办
 * @Param
 * @return
 * @Date 2024/8/19 下午10:23
 **/

func (e *Todos) ListById(userid string) ([]*models.Todos, error) {
	var data []*models.Todos

	err := e.Orm.Model(&models.Todos{}).Where("user_id = ?", userid).Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("该用户暂无待办！")
	}

	return data, nil
}

/**
 * @根据id批量删除待办
 * @Param
 * @return
 * @Date 2024/8/19 下午10:23
 **/

func (e *Todos) Delete(userid string, idStr string, r *ClockRoom) error {
	curDay := time.Now().Format("2006-01-02")
	var data models.Todos
	if len(idStr) == 0 {
		return nil
	}
	curRoom, err := r.GetByUserIdAndDate(userid, curDay)

	if curRoom != nil {
		if strings.Contains(idStr, strconv.Itoa(curRoom.Status)) {
			err = errors.New("存在执行中的待办，请先完成该待办！")
			return err
		}
	}

	// 将字符串形式的 ids 转换成 []int64
	var ids []int64
	for _, id := range strings.Split(idStr, ",") {
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			err = errors.New("数据格式转换错误！")
			return err
		}
		ids = append(ids, idInt)
	}

	result := e.Orm.Model(&data).Where("todo_id IN (?)", ids).Delete(nil)
	if result.Error != nil {
		e.Log.Errorf("db error: %s", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		// 没有找到要删除的记录,认为是正常情况
		return nil
	}
	return nil
}

/**
 * @根据id查询待办
 * @Param
 * @return
 * @Date 2024/8/19 下午10:24
 **/

func (e *Todos) GetById(id int) (*models.Todos, error) {
	var data []*models.Todos

	err := e.Orm.Model(&data).Where("todo_id = ?", id).Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		// 如果出现错误,返回 nil 和错误对象
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("待办不存在！")
	}

	// 返回查询到的 Todos 对象和 nil 错误
	return data[0], nil
}

/**
 * @根据id修改待办
 * @Param
 * @return
 * @Date 2024/8/19 下午10:24
 **/

func (e *Todos) UpdataTodo(c *models.Todos) error {
	var err error
	var data models.Todos
	err = e.Orm.Model(&data).Where("todo_id = ?", c.TodoId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
