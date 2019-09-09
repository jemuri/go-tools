package pool

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_goWorker_run(t *testing.T) {
	tests := []struct {
		name string
		w    *goWorker
	}{
		{
			name: "case1",
			w:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.run()
		})
	}
}

func TestPool_revertWorker(t *testing.T) {
	type args struct {
		worker *goWorker
	}
	tests := []struct {
		name string
		p    *Pool
		args args
		want bool
	}{
		{
			name: "case1",
			p:    nil,
			args: args{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.revertWorker(tt.args.worker); got != tt.want {
				t.Errorf("Pool.revertWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPool(t *testing.T) {
	type args struct {
		size   int
		option []Option
	}
	tests := []struct {
		name string
		args args
		want *Pool
	}{
		{
			name: "case1",
			args: args{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPool(tt.args.size, tt.args.option...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_submit(t *testing.T) {
	type args struct {
		task func()
	}
	tests := []struct {
		name    string
		p       *Pool
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			p:       nil,
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.submit(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Pool.submit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPool_Running(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
		want int
	}{
		{
			name: "case1",
			p:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Running(); got != tt.want {
				t.Errorf("Pool.Running() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_Free(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
		want int
	}{
		{
			name: "case1",
			p:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Free(); got != tt.want {
				t.Errorf("Pool.Free() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_Cap(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
		want int
	}{
		{
			name: "case1",
			p:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Cap(); got != tt.want {
				t.Errorf("Pool.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_retrieveWorker(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
		want *goWorker
	}{
		{
			name: "case1",
			p:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.retrieveWorker(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pool.retrieveWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_incRunning(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
	}{
		{
			name: "case1",
			p:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.incRunning()
		})
	}
}

func TestPool_decRunning(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
	}{
		{
			name: "case1",
			p:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.decRunning()
		})
	}
}

func TestPool_Release(t *testing.T) {
	tests := []struct {
		name string
		p    *Pool
	}{
		{
			name: "case1",
			p:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Release()
		})
	}
}

func Test_Pool(t *testing.T) {
	p := NewPool(3)

	for i := 0; i < 5; i++ {
		f := func() {
			fmt.Printf("即将沉睡%d秒钟",i)
			d := time.Duration(i)
			time.Sleep(time.Second*d)
			fmt.Println()
		}
		selfSubmit(f,p)
	}

	for {
		time.Sleep(time.Second*1)
		fmt.Println("Running: ",p.running)
	}

}

func selfSubmit(f func(),p *Pool) {
	if err := p.submit(f);err != nil {
		fmt.Println("submit出错了: ",err)
		selfSubmit(f,p)
	}
	return
}

func TestName(t *testing.T) {
	f := func() {
		fmt.Printf("即将沉睡%d秒钟",1)
		time.Sleep(time.Second*1)
		fmt.Println()
	}
	fmt.Printf("%v",&f)
}
