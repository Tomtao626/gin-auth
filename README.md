# gin-auth

### 主要实现功能

更多相关项目流程参考[流程图](user.xmind)

1. 建立用户表的增删改查
   - user 用户储存用户的相关信息，不包含敏感信息。(用户名,密码等)
   - user_login 管理用户密码的登录操作,最终关联到 `user` 表查询数据
   - user_oauth 管理用户第三方登录的信息,类似`user_login` (待完善)
2. 用户功能
   - 发送手机验证码 [SendCode]
   - 手机号快速登陆/注册 [PhoneLogin]
   - 账号密码登录 [Login]
   - 账号密码注册 [Register]
   - 获取用户信息 [GetUserInfo]
   - github 登录 [LoginGithub]
3. JWt 认证功能

### 运行步骤:

1. 将 [Sql 文件 user.sql] 导入数据库 `user` 中<br>
   在项目的 conf 配置下。数据库用户、密码默认为 `root` 端口号：`3306`

2. 在项目的根目录下载 swag 命令: `go get -u github.com/swaggo/swag/cmd/swag` <br>

   运行 `swag init` 初始化文档 And 运行 `go run main.go` 跑起来程序

3. 之后在浏览器上运行 http://localhost:8000/swagger/index.html 开始测试功能

### TODO

> 更新 xmind 流程图

1. 其它需求暂定

### 涉及到的知识点(框架)

> 部分功能会展开介绍,如遇问题 Issues

- web 框架 [gin](https://github.com/gin-gonic/gin)
- orm 数据库操作 [gorm](https://github.com/jinzhu/gorm)
- jwt 认证 [jwt-go](https://github.com/dgrijalva/jwt-go)
- 验证 [validator](https://github.com/go-playground/validator)
- 加密 [crypto](https://github.com/golang/crypto)
- 缓存 [redigo](https://github.com/gomodule/redigo)
- 自动生成开发文档
  - [gin-swagger](https://github.com/swaggo/gin-swagger)
  - [swag](https://github.com/swaggo/swag)

### 参考

1. 整体架构入门参考于 [go-gin-example](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)
2. 建表参考于 [廖雪峰的文章](https://www.liaoxuefeng.com/article/001437480923144e567335658cc4015b38a595bb006aa51000)
