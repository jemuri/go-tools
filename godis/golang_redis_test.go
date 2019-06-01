package godis

import (
	"testing"
)

func Test_redis(t *testing.T) {
	p := NewRedisPool()
	err := p.Set("song","zhengjie")
	if err != nil {
		t.Errorf("err: %v",err)
	}
}