package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"GinApi/api/db"
	"GinApi/api/routers"
	"time"
	"os"
	"net/http"
	"log"
)

func main() {
	log.Println("[Server Starting]...")
	// ApiDoc

	// 数据库
	db.InitMySql()
	// 设置开发模式：1、线上环境gin.ReleaseMode。2、开发环境gin.DebugMode
	gin.SetMode(gin.ReleaseMode)
	//获得路由实例
	router := routers.InitRouter()
	// 方法一
	//http.ListenAndServe(":8888", router)
	// 方法二:
	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 监听端口
	server.ListenAndServe()
	log.Println("Server stopped")
	db.CloseDb()
	os.Exit(0)

}
