package strtools

import (
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"strings"
)

//ToString format object to string
func ToString(obj interface{}) string {
	var str string
	if obj == nil {
		return str
	}
	byteArray, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}
	return string(byteArray)
}

// FileSuffix 文件后缀名
func FileSuffix(fileName string) (name, suffix string, err error) {
	if len(fileName) < 4 {
		return "", "", errors.New("名称不合规")
	}

	if !strings.Contains(fileName, ".") {
		return "", "", errors.New("未发现后缀标识")
	}

	index := strings.LastIndex(fileName, ".")
	if index+1 >= len(fileName) {
		return "", "", errors.New("后缀处理发生索引越界")
	}
	name = fileName[:index]
	suffix = fileName[index+1:]
	return name, suffix, nil
}

// generate uuid 32
func UUID() string {
	u := uuid.NewV4()
	return strings.ReplaceAll(u.String(), "-", "")
}

// ChineseConvert 简繁体转换  github.com/stevenyao/go-opencc
//func ChineseConvert(source, patternPath string) (string, error) {
//
//	c := opencc.NewConverter(patternPath)
//	defer c.Close()
//
//	return c.Convert(source), nil
//}