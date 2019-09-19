package strtools

import (
	"testing"
)

func TestFileSuffix(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name       string
		args       args
		wantName   string
		wantSuffix string
		wantErr    bool
	}{
		{
			name: "case1",
			args: args{
				fileName: "中国的.人.png",
			},
			wantName:   "中国的.人",
			wantSuffix: "png",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotSuffix, err := FileSuffix(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSuffix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotName != tt.wantName {
				t.Errorf("FileSuffix() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotSuffix != tt.wantSuffix {
				t.Errorf("FileSuffix() gotSuffix = %v, want %v", gotSuffix, tt.wantSuffix)
			}
		})
	}
}

func TestUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "case1",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UUID(); got != tt.want {
				t.Logf("UUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
