package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

// InitStruct .
func InitStruct(path, confEnv string, confStruct interface{}) {
	if strings.HasSuffix(path, TomlSuffix) {
		conf := os.Getenv(confEnv)
		if conf == "" {
			fmt.Println("Please setting the conf env!!!")
			conf = ConfEnvDefault
		}
		path = strings.Replace(path, TomlSuffix, fmt.Sprintf("_%s%s", conf, TomlSuffix), 1)
	}
	confPath = path
	once.Do(loadConfigToml2(confStruct))
}

func loadConfigToml2(conf interface{}) func() {
	if confPath == "" {
		fmt.Println("loadConfigTomlError: " + ErrEmptyConfPath)
		return nil
	}

	_, err := toml.DecodeFile(confPath, &conf)
	if err != nil {
		fmt.Println(fmt.Sprintf("loadConfigToml-DecodeFileError: %v", err))
		return nil
	}
	fmt.Println("loadConfigToml :", cfg)
	return nil
}
