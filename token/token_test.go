package token

import "testing"

func TestGenerate(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				source: "12345",
			},
			want: "5d65c4b88d3702a6a559208fd7c289de",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.source); got != tt.want {
				t.Logf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
