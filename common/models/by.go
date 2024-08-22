package models

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy int `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy int `json:"updateBy" gorm:"index;comment:更新者"`
	//RoleId   int `json:"roleId" gorm:"index;comment:角色ID"`
	//DeptId   int    `json:"deptId" gorm:"index;comment:部门ID"`
	//Status string `json:"status" gorm:"size:4;comment:状态"`
}

// SetCreateBy 设置创建人id
func (e *ControlBy) SetCreateBy(createBy int) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置修改人id
func (e *ControlBy) SetUpdateBy(updateBy int) {
	e.UpdateBy = updateBy
}

// SetUpdateBy 设置角色id
//func (e *ControlBy) SeRoleId(RoleId int) {
//	e.RoleId = RoleId
//}

// SetUpdateBy 设置部门id
//func (e *ControlBy) SetDeptId(DeptId int) {
//	e.DeptId = DeptId
//}

// SetUpdateBy 设置部门id
//func (e *ControlBy) SetStatus(Status int) {
//	e.Status = strconv.Itoa(Status)
//}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
