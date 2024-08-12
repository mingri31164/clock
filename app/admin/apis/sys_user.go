package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/google/uuid"
	"go-admin/app/admin/common"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/app/admin/utils"
	"go-admin/common/actions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type SysUser struct {
	api.Api
}

// GetPage
// @Summary 列表用户信息数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [get]
// @Security Bearer
func (e SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.SysUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [get]
// @Security Bearer
func (e SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysUser
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserInsertReq true "用户数据"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user [post]
// @Security Bearer
func (e SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update
// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [put]
// @Security Bearer
func (e SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-user/{userId} [delete]
// @Security Bearer
func (e SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
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

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// InsetAvatar
// @Summary 修改头像
// @Description 获取JSON
// @Tags 个人中心
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/avatar [post]
// @Security Bearer
func (e SysUser) InsetAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserAvatarReq{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 数据权限检查
	p := actions.GetPermissionFromContext(c)
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		e.Logger.Debugf("upload avatar file: %s", file.Filename)
		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, filPath)
		if err != nil {
			e.Logger.Errorf("save file error, %s", err.Error())
			e.Error(500, err, "")
			return
		}
	}
	req.UserId = p.UserId
	req.Avatar = "/" + filPath

	err = s.UpdateAvatar(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(filPath, "修改成功")
}

// UpdateStatus 修改用户状态
// @Summary 修改用户状态
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateSysUserStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/status [put]
// @Security Bearer
func (e SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserStatusReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// ResetPwd 重置用户密码
// @Summary 重置用户密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.ResetSysUserPwdReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/reset [put]
// @Security Bearer
func (e SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
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

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.ResetPwd(&req, p)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// UpdatePwd
// @Summary 修改密码
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body dto.PassWord true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/pwd/set [put]
// @Security Bearer
func (e SysUser) UpdatePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.PassWord{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost); err != nil {
		req.NewPassword = string(hash)
	}

	err = s.UpdatePwd(user.GetUserId(c), req.OldPassword, req.NewPassword, p)
	if err != nil {
		e.Logger.Error(err)
		e.Error(http.StatusForbidden, err, "密码修改失败")
		return
	}

	e.OK(nil, "密码修改成功")
}

// GetProfile
// @Summary 获取个人中心用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security Bearer
func (e SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.Id = user.GetUserId(c)

	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = s.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		e.Logger.Errorf("get user profile error, %s", err.Error())
		e.Error(500, err, "获取用户信息失败")
		return
	}
	e.OK(gin.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	}, "查询成功")
}

// GetInfo
// @Summary 获取个人信息
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/getinfo [get]
// @Security Bearer
func (e SysUser) GetInfo(c *gin.Context) {
	req := dto.SysUserById{}
	s := service.SysUser{}
	r := service.SysRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&r.Service).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	p := actions.GetPermissionFromContext(c)
	var roles = make([]string, 1)
	roles[0] = user.GetRoleName(c)
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if user.GetRoleName(c) == "admin" || user.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := r.GetById(user.GetRoleId(c))
		mp["permissions"] = list
		mp["buttons"] = list
	}
	sysUser := models.SysUser{}
	req.Id = user.GetUserId(c)
	err = s.Get(&req, p, &sysUser)
	if err != nil {
		e.Error(http.StatusUnauthorized, err, "登录失败")
		return
	}
	mp["introduction"] = " am a super administrator"
	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if sysUser.Avatar != "" {
		mp["avatar"] = sysUser.Avatar
	}
	mp["userName"] = sysUser.NickName
	mp["userId"] = sysUser.UserId
	mp["deptId"] = sysUser.DeptId
	mp["name"] = sysUser.NickName
	mp["code"] = 200
	e.OK(mp, "")
}

//func Register(ctx *gin.Context) {
//	DB := utils.GetDB()
//	redisdb := utils.InitRedis()
//	//获取参数
//	name := ctx.PostForm("username") //注意字符串要用双引号
//	password := ctx.PostForm("password")
//	email := ctx.PostForm("email")
//	Emailcode := ctx.PostForm("EmailCode")
//
//	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
//		common.ResErr(ctx, "加密错误")
//		return
//	}
//	newUser := models.SysUser{
//		Username: name,
//		Password: string(hasedPassword), //将加密后的密码保存起来
//		Email:    email,
//		DeptId:   1,
//		Status:   strconv.Itoa(2),
//		RoleId:   3,
//	}
//	result, err := redisdb.Get(email).Result()
//	if err != nil {
//		// 处理错误情况
//		common.ResErr(ctx, "请先获取验证码！")
//		return
//	}
//	if result == Emailcode {
//		DB.Create(&newUser)
//		common.ResOK(ctx, "注册成功！")
//		return
//	}
//	if result != Emailcode {
//		common.ResErr(ctx, "验证码错误！")
//		return
//	}
//	common.ResErr(ctx, "注册失败！")
//}

/**
 * 注册接口
 * @param Register
 * @return
 */
func (e SysUser) Register(c *gin.Context) {
	//获取参数
	name := c.PostForm("username") //注意字符串要用双引号
	password := c.PostForm("password")
	email := c.PostForm("email")
	Emailcode := c.PostForm("EmailCode")

	s := service.SysUser{}
	req := dto.Register{}

	req.Username = name
	req.Password = password
	req.Email = email
	req.Emailcode = Emailcode
	req.SeRoleId(3)
	req.SetDeptId(1)
	req.SetStatus(2)

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
	err = s.Register(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	common.ResOK(c, "注册成功！")
}

/**
 * 登录接口
 * @param Login
 * @return
 */
func Login(c *gin.Context) {
	DB := utils.GetDB()
	//获取参数
	username := c.PostForm("name")
	password := c.PostForm("password")

	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		common.ResErr(c, "密码不能少于6位")
		return
	}

	var user models.SysUser
	DB.Where("username = ?", username).First(&user)
	if user.UserId == 0 {
		common.ResErr(c, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		common.ResErr(c, "密码错误")
		return
	}
	common.SetCookie(username, c)
	common.ResOK(c, "登录成功")
}

func Start(c *gin.Context) {
	uname, _ := c.Get("user_name")
	db := utils.GetDB()
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	// 查询当前用户当天是否有打卡记录，有则更新，无则新增
	var duration models.Durations
	db.Where("name = ? AND date = ?", uname, currentDate).First(&duration)
	// 开始时间不为空，说明正在计时中
	if duration.StartTime != "" {
		common.ResOK(c, "已完成开始打卡")
		return
	}
	// 开始时间为空则判断今日是否打过卡，打过卡说明已经创建过记录，把记录更新即可
	if duration.ID != 0 {
		db.Model(&duration).Update("start_time", currentTime)
		common.ResOK(c, "开始打卡成功")
		return
	}
	// 未打过卡就新增本日记录
	newDur := models.Durations{
		Name:      uname.(string),
		Date:      currentDate,
		StartTime: currentTime,
		Dur:       0,
	}
	db.Create(&newDur)
	common.ResOK(c, "开始打卡成功")
}

func End(c *gin.Context) {
	uname, _ := c.Get("user_name")
	db := utils.GetDB()
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	// 查询当前用户当天是否有打卡记录
	var duration models.Durations
	db.Where("name = ? AND date = ?", uname, currentDate).First(&duration)
	// 提示未进行开始打卡
	if duration.StartTime == "" {
		common.ResErr(c, "未进行开始打卡")
		return
	}
	// 结算本次打卡时间
	t1, err1 := time.Parse("15:04:05", duration.StartTime)
	t2, err2 := time.Parse("15:04:05", currentTime)
	if err1 != nil || err2 != nil {
		common.ResErr(c, "时间解析出错")
		return
	}
	durObj := t2.Sub(t1)
	dur := int(durObj.Seconds())
	todayDur := duration.Dur + dur
	// 更新本日记录
	db.Model(&duration).Updates(map[string]interface{}{
		"StartTime": "",
		"Dur":       todayDur,
	})
	// 返回具体时分秒
	h := dur / 3600
	m := (dur % 3600) / 60
	s := dur % 60
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "结束打卡成功",
		"hours":   h,
		"minutes": m,
		"seconds": s,
	})
}
