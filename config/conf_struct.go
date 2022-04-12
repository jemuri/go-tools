package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"strings"
)

// InitStruct .
func InitStruct(path, env string, confStruct interface{}) {
	if strings.HasSuffix(path, TomlSuffix) {
		conf := os.Getenv(env)
		if conf == "" {
			fmt.Println("Please setting the conf env!!!")
			conf = ConfEnvDefault
		}
		path = strings.Replace(path, TomlSuffix, fmt.Sprintf("_%s%s", conf, TomlSuffix), 1)
	}
	confPath = path
	loadConfigToml2(confStruct)
}

func loadConfigToml2(confStruct interface{}) {
	if confPath == "" {
		fmt.Println("loadConfigTomlError: " + ErrEmptyConfPath)
		return
	}

	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(fmt.Sprintf("loadConfigToml-ReadFileError: %v", err))
	}
	_, err = toml.Decode(string(data), confStruct)
	if err != nil {
		fmt.Println(fmt.Sprintf("loadConfigToml-toml.DecodeError: %v", err))
	}

	fmt.Println("confStruct :", confStruct)
}
