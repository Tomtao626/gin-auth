package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sun-wenming/gin-auth/models"
	"github.com/sun-wenming/gin-auth/pkg/gredis"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/oauth"
	"github.com/sun-wenming/gin-auth/pkg/setting"
	"github.com/sun-wenming/gin-auth/routers"
	"net/http"
)

// @title 用户登录注册/认证示例
// @description 用户登录注册/认证示例 Golang语言编写 Gin框架
// @termsOfService https://github.com/sun-wenming/gin-auth

// @license.name Apache License 2.0
// @license.url https://github.com/sun-wenming/gin-auth/blob/master/LICENSE

// @host localhost:8000
// @BasePath

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name jwtToken
func main() {
	setting.Setup()
	logging.Setup()
	models.Setup()
	// 缓存 redis
	err := gredis.Setup()
	if err != nil {
		logging.GetLogger().Fatalln(err)
	}
	oauth.Setup()

	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// endless.NewServer(endPoint,router)
	// server.BeforeBegin = func(add string) {
	//     log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	//     log.Printf("Server err: %v", err)
	// }
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
