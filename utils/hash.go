package utils

import (
	"crypto/md5"
	"fmt"
)

const salt = "jimmy loves big meetting room "

func Md5(data string) string {
	data = data + salt
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}
