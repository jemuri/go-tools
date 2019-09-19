package arr

import "testing"

func TestContain(t *testing.T) {
	type args struct {
		value interface{}
		arr   interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				value: 4,
				arr:   []int{1,2,3,4,5,6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.value, tt.args.arr); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}
