package db

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"database/sql"
)

// 1、原生---数据库连接初始化
var DB *sql.DB
func InitMySql() {
	log.Println("[init database]......")
	mysqlConfig := "root:Qwe110120*@/go_schema?charset=utf8"
	var err error
	DB, err = sql.Open("mysql", mysqlConfig)
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
