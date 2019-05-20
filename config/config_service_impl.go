package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/PaesslerAG/jsonpath"
	"os"
	"strings"
	"sync"
)

const (
	GuaZiEnv   = "GUAZI_ENV" //Attention
	TomlSuffix = ".toml"
	ErrEmptyConfPath = "config path is null!"
)

var (
	cfg      map[string]interface{}
	once     sync.Once
	confPath string
)

// Init .
func Init(path string) {
	if strings.HasSuffix(path, TomlSuffix) {
		path = strings.Replace(path, TomlSuffix, fmt.Sprintf("_%s%s", os.Getenv(GuaZiEnv), TomlSuffix), 1)
	}
	confPath = path
	once.Do(loadConfigToml)
}

// Get .
func Get(key string) (interface{}, error) {
	once.Do(loadConfigToml)

	value, err := jsonpath.Get(convert(key), cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("jsonpath.Get happened error: %v", err))
	}

	return value, nil
}

func convert(key string) string {
	if !strings.HasPrefix(key, "$.") {
		key = "$." + key
	}
	return strings.Replace(key, "/", ".", -1)
}

func loadConfigToml() {
	if confPath == "" {
		fmt.Println(ErrEmptyConfPath)
		panic(ErrEmptyConfPath)
	}
	_, err := toml.DecodeFile(confPath, &cfg)
	if err != nil {
		fmt.Println(fmt.Sprintf("DecodeFileError: %v", err))
		panic(err)
	}
	fmt.Println("loadConfigToml :", cfg)
}
