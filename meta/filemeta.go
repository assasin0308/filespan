package meta

// 定义文件原信息
type FileMeta struct {
	FileSha1 string // 文件加密
	FileName string // 文件名称
	FileSize int64 // 文件大小
	Location string // 文件存储位置
	UploadTime string // 文件上传时间
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}
// 新增&更新文件元信息
func UpdateFileMeta(filemeta FileMeta) {
	fileMetas[filemeta.FileSha1] = filemeta
}
// 通过sha1获取文件元信息
func GetFileMeta(filesha1 string) FileMeta {
	return fileMetas[filesha1]
}
// 批量获取文件元信息列表
func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta,len(fileMetas))
	for _,v := range fileMetas {
		fMetaArray = append(fMetaArray,v)
	}
	//sort.Sort(ByUp)
	return fMetaArray[0:count]
}

// 删除元信息
func RemoveFileMeta(filesha1 string) {
	delete(fileMetas,filesha1)
}