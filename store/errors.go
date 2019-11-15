package store

import "errors"

var ErrReadFile = errors.New("读取账号文件错误")
var ErrWriteFile = errors.New("写入账号文件错误")
var ErrParse = errors.New("序列化账号错误")
