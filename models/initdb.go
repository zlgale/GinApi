package models

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"database/sql"
)

// 1、原生---数据库连接初始化
var DB *sql.DB
func Init() {
	log.Println("[init database]......")
	mysql_config := "root:Qwe110120*@/go_schema?charset=utf8"
	var err error
	DB, err = sql.Open("mysql", mysql_config)
	if err != nil {
		log.Fatalf("Fail to open mysql: %v", err)
	}
	// 连接池的空闲大小
	DB.SetMaxIdleConns(1000)
	// 最大打开连接数
	DB.SetMaxOpenConns(2000)
	// 日志
	os.Create("mysql.log")
	// 测试是否连接成功
	if err := DB.Ping(); err != nil {
		log.Println("DB.Ping():", DB.Ping())
		return
	}
	log.Println("[end init database]......")
}
func CloseDb(){
	log.Println("DB.Ping():", DB.Ping())
	defer DB.Close()
}

//// 2、Xorm---数据库连接初始化
//var X *xorm.Engine
//func Init() {
//	log.Println("[init database]......")
//	mysql_config := "root:Qwe110120*@/go_schema?charset=utf8"
//	var err error
//	X, err = xorm.NewEngine("mysql", mysql_config)
//	if err != nil {
//		log.Fatalf("Fail to create engine: %v", err)
//	}
//	// 连接池的空闲大小
//	X.SetMaxIdleConns(240)
//	// 最大打开连接数
//	X.SetMaxOpenConns(240)
//	// 日志
//	f, err := os.Create("sql.log")
//	if err != nil {
//		println(err.Error())
//		return
//	}
//	X.SetLogger(xorm.NewSimpleLogger(f))
//	X.ShowSQL(true)
//	//x.Logger().SetLevel(core.LOG_DEBUG)
//	// 测试是否连接成功
//	if err := X.Ping(); err != nil {
//		log.Println("engin.Ping():", X.Ping())
//		return
//	}
//	// 前缀映射
//	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "prefix_")
//	X.SetTableMapper(tbMapper)
//	//err = X.Sync2(new(user.Account), new(item.Item))
//	//if err != nil {
//	//	log.Fatalf("Fail to sync2 database: %v", err)
//	//}
//	log.Println("[end init database]......")
//}
