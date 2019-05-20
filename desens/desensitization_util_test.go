package desens

import "testing"

func TestPhoneNumber(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				phone: "234",
			},
			want: "234",
		},{
			name: "case2",
			args: args{
				phone: "18566667777",
			},
			want: "185****7777",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PhoneNumber(tt.args.phone); got != tt.want {
				t.Errorf("PhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
