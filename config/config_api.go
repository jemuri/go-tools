package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	Empty      = ""
	Endpoint   = "endpoint"
	AppKey     = "app_key"
	Secret     = "secret"
	TomlSuffix = ".toml"
)

// Init .
func Init(path, confEnv string) {
	if strings.HasSuffix(path, TomlSuffix) {
		conf := os.Getenv(confEnv)
		if conf == "" {
			panic("Please setting the conf env!!!")
		}
		path = strings.Replace(path, TomlSuffix, fmt.Sprintf("_%s%s", conf, TomlSuffix), 1)
	}
	confPath = path
	once.Do(loadConfigToml)
}

// CertainString 确认必定返回有效string
func CertainString(key string) string {
	value, err := Get(key)
	if err != nil {
		fmt.Println("LocalConfig: CertainString convert string err: ", err.Error())
		return ""
	}

	val, ok := value.(string)
	if !ok {
		return "LocalConfig: CertainString convert string fail"
	}
	return val
}

// CertainInt64 确认必定返回有效int64
func CertainInt64(key string) int64 {
	value, err := Get(key)
	if err != nil {
		fmt.Println("LocalConfig: CertainInt64 convert int64 err: ", err.Error())
		return -1
	}

	val, ok := value.(int64)
	if !ok {
		fmt.Println("LocalConfig: CertainInt64 convert int64 fail")
		return -1
	}
	return val
}

// CertainSign 获取endpoint及验签
func CertainSign(key string) (endpoint, appKey, secret string) {
	value, err := Get(key)
	if err != nil {
		fmt.Println("LocalConfig: CertainSign err: ", err.Error())
		return
	}

	val, ok := value.(map[string]interface{})
	if !ok {
		fmt.Println("LocalConfig: CertainSign map err: ", err.Error())
		return
	}

	uri, ok := val[Endpoint].(string)
	if !ok {
		fmt.Println("LocalConfig: CertainSign convert Endpoint fail")
		return
	}

	app, ok := val[AppKey].(string)
	if !ok {
		fmt.Println("LocalConfig: CertainSign convert AppKey fail")
		return
	}

	sec, ok := val[Secret].(string)
	if !ok {
		fmt.Println("LocalConfig: CertainSign convert Secret fail")
		return
	}

	return uri, app, sec
}
