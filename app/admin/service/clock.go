package service

/**
 * @Description
 * @Author mingri
 * @Date 2024/8/19 下午7:33
 **/

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"strconv"
	"strings"
	"time"
)

type Clock struct {
	service.Service
}

/**
 * @新增打卡记录（打卡逻辑）
 * @Param
 * @return
 * @Date 2024/8/19 下午10:19
 **/

func (e Clock) Insert(c *dto.AddClock, t *Todos, clock *Clock, u *SysUser, r *ClockRoom) (int, error) {
	var err error
	var data models.Clock
	var newRoom models.ClockRoom
	curTime := time.Now().Format("2006-01-02")

	if c.Place == "" {
		err = errors.New("请选择打卡地点！")
		return -1, err
	}

	todo, err := t.GetById(c.TodoId)
	if todo == nil {
		err = errors.New("待办不存在！")
		return -1, err
	}
	if todo.Status == -1 {
		err = errors.New("待办执行中，无法打卡！")
		return todo.TodoId, err
	}

	//查询当前用户当天是否加入自习室
	curRoom, err := r.GetByUserIdAndDate(strconv.Itoa(c.UserID), curTime)
	//已加入自习室
	if curRoom != nil {
		//查看当前用户是否还在打卡状态
		if curRoom.Status != -1 {
			err = errors.New("正在打卡中，请先完成当前打卡！")
			return curRoom.Status, err
		}
		//如果当前打卡地点不在今日自习地点中，则添加
		if !strings.Contains(curRoom.Place, c.Place) {
			if len(curRoom.Place) > 0 {
				curRoom.Place = fmt.Sprintf("%s,%s", curRoom.Place, c.Place)
			} else {
				curRoom.Place = c.Place
			}
		}
	}

	//未加入自习室
	//查询昨天是否有打卡记录
	yesterday := time.Now().AddDate(0, 0, -1)
	yesterdayClock, err := clock.ListByDate(yesterday)
	todayClock, err := clock.ListByDate(time.Now())

	//查询当前用户信息
	user, err := u.GetByUserId(strconv.Itoa(c.UserID))
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return -1, err
	}

	if len(todayClock) == 0 {
		//今天未打卡，则总打卡天数 + 1
		user.Sum += 1
		if curRoom == nil {
			newRoom.UserID = c.UserID
			newRoom.Place = c.Place
			newRoom.Status = c.TodoId
			err = r.Insert(&newRoom)
			if err != nil {
				e.Log.Errorf("db error: %s", err)
				return -1, err
			}
		}
		if curRoom != nil {
			curRoom.Status = c.TodoId
			curRoom.Place = c.Place
		}
		//如果昨天有打卡记录且今天还未打卡，则连续打卡天数 + 1
		if len(yesterdayClock) != 0 {
			user.Continuous += 1
		}
		//如果昨天没有打卡记录且今天还未打卡，则连续打卡天数重置为1
		if len(yesterdayClock) == 0 {
			user.Continuous = 1
		}

		err = u.UpdateUser(user)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return -1, err
		}
	}
	if len(todayClock) != 0 {
		//今天已有打卡，则更新当前打卡事项status
		curRoom.Status = c.TodoId
		err = r.UpdataCur(curRoom)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return -1, err
		}
	}

	todo.Status = -1
	err = t.UpdataTodo(todo)
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return -1, err
	}

	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return -1, err
	}
	return data.ClockId, nil
}

/**
 * @删除打卡记录
 * @Param
 * @return
 * @Date 2024/8/19 下午10:20
 **/

func (e *Clock) Delete(idsStr string) error {
	var data models.Clock
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
		clock, _ := e.GetById(strconv.FormatInt(id, 10))

		if clock != nil {
			if clock.ClockTime == 0 {
				return errors.New("存在未结束的打卡任务！")
			}
		}

	}

	result := e.Orm.Model(&data).Where("clock_id IN (?)", ids).Delete(nil)
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
 * @根据年月日查询打卡记录
 * @Param
 * @return
 * @Date 2024/8/19 下午10:20
 **/

func (e *Clock) ListByDate(clockDate time.Time) ([]*models.Clock, error) {
	var data []*models.Clock

	// 格式化clockDate，只取年月日部分
	formattedClockDate := clockDate.Format("2006-01-02")

	// 执行查询，比较年月日部分
	err := e.Orm.Where("DATE(start_at) = ?", formattedClockDate).Find(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("打卡记录不存在！")
	}
	return data, nil
}

/**
 * @结束打卡
 * @Param
 * @return
 * @Date 2024/8/20 下午4:38
 **/

func (e Clock) EndClock(todoId int, u *SysUser, t *Todos, r *ClockRoom) error {
	curTime := time.Now().Format("2006-01-02")

	data, err := e.GetByTodoId(todoId)
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	//查看该打卡记录是否已结束
	if data.ClockTime == 0 {
		user, err := u.GetByUserId(strconv.Itoa(data.UserID))
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
		todo, err := t.GetById(todoId)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
		if todo.Status == 2 {
			return errors.New("待办已结束，不能重复打卡！")
		}
		curRoom, err := r.GetByUserIdAndDate(strconv.Itoa(data.UserID), curTime)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		//数据回显更新
		endTime := time.Now()
		duration := endTime.Sub(data.StartAt)
		minutes := int(duration.Minutes())
		if minutes < 1 {
			data.EndAt = endTime
			data.ClockTime = -1
			err = e.UpdataClock(data)
			if err != nil {
				e.Log.Errorf("db error: %s", err)
				return err
			}
			err := e.Delete(strconv.Itoa(data.ClockId))
			if err != nil {
				e.Log.Errorf("db error: %s", err)
				return err
			}
			curRoom.Status = -1
			err = r.UpdataCur(curRoom)
			if err != nil {
				e.Log.Errorf("db error: %s", err)
				return err
			}
			todo.Status = 2
			err = t.UpdataTodo(todo)
			if err != nil {
				e.Log.Errorf("db error: %s", err)
				return err
			}

			if curRoom.TodoIds == "" {
				//如果这是今日第一次加入自习室，则删除当前自习室记录
				err = r.DeleteRoom(strconv.Itoa(curRoom.RoomId))
				if err != nil {
					e.Log.Errorf("db error: %s", err)
					return err
				}
			}
			return errors.New("打卡时长少于一分钟,不计入时长！")
		}
		// 将 duration 转换为时分秒格式
		//hours := int(duration.Hours())
		//minutes := int(duration.Minutes()) % 60
		//seconds := int(duration.Seconds()) % 60
		//clockTime := fmt.Sprintf("%02d时%02d分%02d秒", hours, minutes, seconds) //将duration转换为字符串

		//更新当前打卡记录数据
		data.EndAt = endTime
		data.ClockTime = minutes

		//更新当前用户打卡总时长
		user.TimeTotal += minutes

		//更新待办执行时长和次数
		todo.Status = 2
		todo.Duration += minutes
		todo.Loop += 1

		//更细自习室数据
		curRoom.ClockTime += minutes
		curRoom.Status = -1
		curTodos := curRoom.TodoIds
		curTodo := strconv.Itoa(todo.TodoId)
		//如果当前待办不在自习室待办列表中，则添加
		if !strings.Contains(curTodos, curTodo) {
			if len(curTodos) > 0 {
				curRoom.TodoIds = fmt.Sprintf("%s,%s", curTodos, curTodo)
			} else {
				curRoom.TodoIds = curTodo
			}
		}

		//数据更新
		err = e.UpdataClock(data)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		err = r.UpdataCur(curRoom)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		err = u.UpdateUser(user)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}

		err = t.UpdataTodo(todo)
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return err
		}
		return nil
	}
	err = errors.New("不能重复当前打卡！")
	return err
}

/**
 * @根据用户id查询所有打卡记录
 * @Param
 * @return
 * @Date 2024/8/20 下午4:38
 **/

func (e Clock) ListByUserId(userid string) ([]*models.Clock, error) {
	var err error
	var ClockList []*models.Clock

	err = e.Orm.Where("user_id = ?", userid).Find(&ClockList).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	return ClockList, nil
}

func (e Clock) GetById(clockid string) (*models.Clock, error) {
	var err error
	var clock []*models.Clock

	err = e.Orm.Model(&clock).Where("clock_id = ?", clockid).Find(&clock).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(clock) == 0 {
		return nil, errors.New("打卡记录不存在！")
	}
	return clock[0], nil
}

func (e Clock) GetByTodoId(todoId int) (*models.Clock, error) {
	var err error
	var clock []*models.Clock
	today := time.Now().Format("2006-01-02")

	err = e.Orm.Model(&clock).Where("todo_id = ? AND DATE(start_at) = ? AND clock_time = 0", todoId, today).
		Find(&clock).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if len(clock) == 0 {
		return nil, errors.New("打卡记录不存在！")
	}
	return clock[0], nil
}

func (e Clock) UpdataClock(c *models.Clock) error {
	var err error
	var data models.Clock

	err = e.Orm.Model(&data).Where("clock_id = ?", c.ClockId).Updates(c).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}
