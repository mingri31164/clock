package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
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
	err = e.Orm.Model(&data).Where("id = ?", c.Id).Updates(c).Error
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
	// 执行查询，比较年月日部分
	err := e.Orm.Where("date = ?", date).Find(&data).Error
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
		curRoom, err := e.GetById(strconv.FormatInt(id, 10))
		if err != nil {
			return err
		}
		if curRoom.Status != -1 {
			return errors.New("存在未结束的打卡任务！")
		}
	}

	result := e.Orm.Model(&data).Where("id IN (?)", ids).Delete(nil)
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

	err = e.Orm.Model(&data).Where("id = ?", roomid).Find(&data).Error
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
