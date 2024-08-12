# clock-main

  <img align="right" width="320" src="http://116.205.189.126:9000/clock-bucket/f1a2ee1f-d041-4dfa-8225-609ddc41d8b3.jpg">



基于Gin + Vue + Element UI 前后端分离的打卡考勤项目后台管理系统

[在线文档](https://www.go-admin.pro)

[前端项目](https://github.com/go-admin-team/go-admin-ui)


## 🎬 在线体验

后台管理页面：[https://vue2.clock-main.dev](http://116.205.189.126:8090/)
> ⚠️⚠️⚠️ 账号 / 密码： admin / 123456

## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库

- 配置文件简单的模型映射，快速能够得到想要的配置

- 代码生成工具

- 表单构建工具

- 多指令模式

- 多租户的支持

- TODO: 单元测试

## 🎁 内置

1. 多租户：系统默认支持多租户，按库分离，一个库一个租户。
1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识，接口权限等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 参数管理：对系统动态配置常用参数。
8. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
9. 登录日志：系统登录日志记录查询包含登录异常。
1. 接口文档：根据业务代码自动生成相关的api接口文档。
1. 代码生成：根据数据表结构生成对应的增删改查相对应业务，全程可视化操作，让基本业务可以零代码实现。
1. 表单构建：自定义页面样式，拖拉拽实现页面布局。
1. 服务监控：查看一些服务器的基本信息。
1. 内容管理：demo功能，下设分类管理、内容管理。可以参考使用方便快速入门。
1. 定时任务：自动化任务，目前支持接口调用和函数调用。



## 📦 本地开发

### 环境要求

go 1.21

node版本: v18.19.1

npm版本: 10.2.4



### 启动说明

> 重点注意：两个项目必须放在同一文件夹下

#### 服务端启动说明

```bash
# 进入 clock-main 后端项目
cd ./clock-main

# 更新整理依赖
go mod tidy

# 编译项目
go build

# 修改配置 
# 文件路径  clock-main/config/settings.yml

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径（非必须）
```



:::

#### 初始化数据库，以及服务启动

``` bash
# 首次配置需要初始化数据库资源信息
# macOS or linux 下使用
go run main.go migrate -c config/settings.yml

# ⚠️注意:windows 下使用
go run main.go migrate -c config\settings.yml


# 启动项目，也可以用IDE进行调试
# macOS or linux 下使用
go run main.go server -c config/settings.yml


# ⚠️注意:windows 下使用
go run main.go server -c config\settings.yml

```

#### sys_api 表的数据如何添加（非必须）

在项目启动时，使用`-a true` 系统会自动添加缺少的接口数据
```bash
go run main.go server -c config\settings.dev.yml -a true

```


#### 文档生成

```bash
go generate
```


### UI交互端启动说明

```bash
# 安装依赖
npm install

# 建议使用镜像安装并忽略eslint版本冲突问题
npm install --registry=https://registry.npmmirror.com --legacy-peer-deps

# 启动服务
npm run dev

# 打包前端项目
npm run build:prod
```

## 💎 开发人员


<span style="margin: 0 5px;" ><a href="https://github.com/mingri31164" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/150683291?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/d2bz" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/139675120?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>

