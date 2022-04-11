package config

import (
	"fmt"
	"testing"
)

var ConfigStruct = struct {
	AppName string `json:"app_name"`
	DbDsn   string `json:"db_dsn" toml:"db_dsn"`
	Admin   struct {
		Port string `json:"port"`
	}
	Server struct {
		Port string
	}
	Database struct {
		Ports []int64 `json:"ports"`
	}
}{}

func TestInitStruct(t *testing.T) {
	type args struct {
		path    string
		confEnv string
		conf    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				path:    "config/config.toml",
				confEnv: "dev",
				conf:    &ConfigStruct,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitStruct(tt.args.path, tt.args.confEnv, tt.args.conf)
			fmt.Println(tt.args.conf)
		})
	}
}
