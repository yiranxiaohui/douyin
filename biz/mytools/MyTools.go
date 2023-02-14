package mytools

import (
	"crypto/md5"
	"fmt"
)

func GetMD5String(s string) (result string) {
	bytes := []byte(s)
	res:=md5.Sum(bytes) //返回值：[Size]byte 数组
	result=fmt.Sprintf("%x",res)   //通过fmt.Sprintf()方法格式化数据
	return
}
