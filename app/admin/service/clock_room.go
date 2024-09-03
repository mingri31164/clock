package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/vo"
	"strconv"
	"strings"
	"time"
)

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/20 下午7:04
 **/

type ClockRoom struct {
	service.Service
}

/**
 * @加入自习室
 * @Param
 * @return
 * @Date 2024/8/20 下午11:49
 **/

func (e *ClockRoom) Insert(c *models.ClockRoom) error {
	var err error
	var data models.ClockRoom
	var user models.SysUser
	var dept models.SysDept
	var i int64
	c.Date = time.Now().Format("2006-01-02")
	err = e.Orm.Model(&data).Where("date = ? AND user_id = ?", c.Date, c.UserID).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("今日已加入自习室！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	err = e.Orm.Model(&user).Where("user_id = ?", c.UserID).Find(&user).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	err = e.Orm.Model(&dept).Where("dept_id = ?", user.DeptId).Find(&dept).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.Username = user.Username
	c.Dept = dept.DeptName
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

/**
 * @更新自习室状态
 * @Param
 * @return
 * @Date 2024/8/20 下午11:52
 **/

func (e *ClockRoom) UpdataCur(c *models.ClockRoom) error {
	var err error
	var data models.ClockRoom
	err = e.Orm.Model(&data).Where("room_id = ?", c.RoomId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

/**
 * @查询今日自习室状态
 * @Param
 * @return
 * @Date 2024/8/20 下午11:54
 **/

func (e *ClockRoom) ListRoom(date string) ([]*models.ClockRoom, error) {
	var data []*models.ClockRoom
	// 执行查询，比较年月日部分，并按 clock_time 字段倒序排序
	err := e.Orm.Where("date = ?", date).Order("clock_time DESC").Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("未找到当前自习信息！")
	}
	return data, nil
}

/**
 * @根据用户id查看自习室情况
 * @Param
 * @return
 * @Date 2024/8/21 上午12:31
 **/

func (e *ClockRoom) GetByUserIdAndDate(userid string, date string) (*models.ClockRoom, error) {
	var data []*models.ClockRoom
	// 执行查询，比较年月日部分
	err := e.Orm.Where("user_id = ? AND date = ?", userid, date).Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("未找到当前自习信息！")
	}
	return data[0], nil
}

/**
 * @删除打卡记录
 * @Param
 * @return
 * @Date 2024/8/22 下午5:38
 **/

func (e *ClockRoom) DeleteRoom(idsStr string) error {
	var data models.ClockRoom
	if len(idsStr) == 0 {
		return nil
	}
	// 将字符串形式的 ids 转换成 []int64
	var ids []int64
	for _, id := range strings.Split(idsStr, ",") {
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errors.New("数据格式错误！")
		}
		ids = append(ids, idInt)
	}
	for _, id := range ids {
		// 对 id 执行操作
		curRoom, _ := e.GetById(strconv.FormatInt(id, 10))
		if curRoom != nil {
			if curRoom.Status != -1 {
				return errors.New("存在未结束的打卡任务！")
			}
		}
	}

	result := e.Orm.Model(&data).Where("room_id IN (?)", ids).Delete(nil)
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
 * @根据id获取打卡记录
 * @Param
 * @return
 * @Date 2024/8/22 下午5:39
 **/

func (e ClockRoom) GetById(roomid string) (*models.ClockRoom, error) {
	var err error
	var data []*models.ClockRoom

	err = e.Orm.Model(&data).Where("room_id = ?", roomid).Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("未找到当前自习信息！")
	}
	return data[0], nil
}

/**
 * @根据用户id获取打卡记录
 * @Param
 * @return
 * @Date 2024/8/22 下午5:40
 **/

func (e *ClockRoom) GetByUserId(userid string) ([]*models.ClockRoom, error) {
	var data []*models.ClockRoom
	err := e.Orm.Where("user_id = ?", userid).Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("未找到当前自习信息！")
	}
	return data, nil
}

func (e *ClockRoom) ListFinishTodes(roomid string) (*vo.FinishTodos, error) {
	var finishTodos vo.FinishTodos
	var clock models.Clock
	var todos models.Todos

	curRoom, err := e.GetById(roomid)
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	// Convert string of ids to []int64
	var todoIds []int64
	for _, id := range strings.Split(curRoom.TodoIds, ",") {
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, errors.New("数据格式错误！")
		}
		todoIds = append(todoIds, idInt)
	}

	var todoNames []string
	var clockTimes []string
	// 格式化 curRoom.Date 为当天的开始时间
	startDate := vo.FormatToStartOfDay(curRoom.Date)

	err = e.Orm.Model(&todos).Select("name").
		Where("todo_id in (?)", todoIds).
		Pluck("name", &todoNames).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	err = e.Orm.Model(&clock).Select("clock_time").
		Where("todo_id IN (?) AND end_at >= ? AND end_at < ?",
			todoIds, startDate, startDate.AddDate(0, 0, 1)).
		Pluck("clock_time", &clockTimes).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}

	finishTodos.Todoname = strings.Join(todoNames, ",")
	finishTodos.Clocktime = strings.Join(clockTimes, ",")

	return &finishTodos, nil
}
