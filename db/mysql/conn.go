package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

var db *sql.DB

func init() {
	db,_ := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/filespan?charset=utf8")
	db.SetMaxIdleConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect Mysql,error:",err.Error())
		os.Exit(1)
	}
}

// 返回数据库连接对象
func DBConnect() *sql.DB {
	return db
}
