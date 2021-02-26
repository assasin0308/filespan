package db

import (
	mydb "filespan/db/mysql"
	"fmt"
)

// 用户注册
func UserSign(username string,passwd string) bool {
	stmt,err := mydb.DBConnect().Prepare(
		"INSERT ignore INTO tbl_user(`user_name`,`user_pwd`) VALUES(?,?) ")
	if err != nil {
		fmt.Println("1111111111")
		fmt.Printf("Failed to insert,error:",err.Error())
		return false
	}
	defer stmt.Close()
	ret,err := stmt.Exec(username,passwd)
	if err != nil {
		fmt.Println("failed insert,err" + err.Error())
		return  false
	}
	if rowsAffected,err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}