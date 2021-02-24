package db

import (
	"database/sql"
	mydb "filespan/db/mysql"
	"fmt"
)
// 上传文件完成,保存meta信息
func OnFileUploadFinished(filehash,filename,fileaddr string,filesize int64) bool {
	stmt,err := mydb.DBConnect().Prepare("INSERT INTO tbl_file " +
		"(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`)" +
		" VALUES (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement,error:",err.Error())
		return false
	}
	defer stmt.Close()
	result,err := stmt.Exec(filehash,filename,fileaddr,filesize)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rowsaffected,err := result.RowsAffected();nil == err {
		if rowsaffected <= 0 {
			fmt.Printf("Failed with hash:%s has been uploaded before",filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize int64
	FileAddr sql.NullString
}
// 根据文件hash获取文件元信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt,err := mydb.DBConnect().Prepare("SELECT file_sha1,file_name,file_addr,file_size from tbl_file " +
		"WHERE file_sha1 = ? AND status = 1 LIMIT 1 ")
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	defer stmt.Close()
	table_file := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&table_file.FileHash,&table_file.FileName,&table_file.FileSize,&table_file.FileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil,err
	}
	return &table_file,nil
}
