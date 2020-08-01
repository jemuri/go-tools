package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/PaesslerAG/jsonpath"
	"strings"
	"sync"
)

const (
	ErrEmptyConfPath = "config path is null!"
)

var (
	cfg      map[string]interface{}
	once     sync.Once
	confPath string
)

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
		fmt.Println("loadConfigTomlError: " + ErrEmptyConfPath)
		return
	}
	_, err := toml.DecodeFile(confPath, &cfg)
	if err != nil {
		fmt.Println(fmt.Sprintf("loadConfigToml-DecodeFileError: %v", err))
		return
	}
	fmt.Println("loadConfigToml :", cfg)
}
