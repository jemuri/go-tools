package godis

import (
	"fmt"
	"github.com/jemuri/go-tools/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Init("../../conf/config.toml", "CONF_ENV")
	os.Exit(m.Run())
}

type Person struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

func Test_redis(t *testing.T) {
	p := NewRedisPool()

	a := &Person{
		ID:   18,
		Name: "song",
	}
	err := p.Set("song",a)
	if err != nil {
		t.Errorf("err: %v",err)
	}
	aa,err :=p.Get("song")
	if err != nil {
		t.Errorf("get error: %v",err)
	}
	fmt.Print(aa)
}