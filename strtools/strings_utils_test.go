package strtools

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
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

const fan = `{
    "author": "胡宿", 
    "paragraphs": [
      "五粒青松護翠苔，石門岑寂斷纖埃。", 
      "水浮花片知仙路，風遞鸞聲認嘯臺。", 
      "桐井曉寒千乳斂，茗園春嫩一旗開。", 
      "馳煙未勒山亭字，可是英靈許再來。"
    ], 
    "strains": [
      "仄仄平平仄仄平，仄平平仄仄平平。", 
      "仄平平仄平平仄，平仄平平仄仄平。", 
      "平仄仄平平仄仄，仄平平仄仄平平。", 
      "平平仄仄平平仄，仄仄平平仄仄平。"
    ], 
    "title": "沖虛觀"
  }, 
  {
    "author": "胡宿", 
    "paragraphs": [
      "天臺封詔紫泥馨，馬首前瞻北斗城。", 
      "人在函關先望氣，帝于京兆最知名。", 
      "一區東第趨晨近，數刻西廂接晝榮。", 
      "正是兩宮裁化日，百金雙璧拜虞卿。"
    ], 
    "strains": [
      "平平仄仄仄平平，仄仄平平仄仄平。", 
      "平仄平平平仄仄，仄平平仄仄平平。", 
      "仄平平仄平平仄，仄仄平平仄仄平。", 
      "仄仄仄平平仄仄，仄平平仄仄平平。"
    ], 
    "title": "淮南發運趙邢州被詔歸闕"
  }, `

func TestChineseConvert(t *testing.T) {
	type args struct {
		source      string
		patternPath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				source:      fan,
				patternPath: "/usr/local/Cellar/opencc/1.0.5/share/opencc/tw2sp.json",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Open("/Users/songsong/dev/WebSpace/chinese-poetry-master/json/poet.song.0.json")
			if err != nil {
				t.Logf("出错了，Open: %v", err)
			}
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				t.Logf("出错了，Open: %v", err)
			}

			fmt.Print("打印一个转过前的结果: ", string(bytes))
			fmt.Println()
			got, err := ChineseConvert(string(bytes), tt.args.patternPath)
			fmt.Print("打印一个转过结果: ", got)
			fmt.Println()
			if (err != nil) != tt.wantErr {
				//t.Logf("ChineseConvert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_unit(t *testing.T) {
	s := `abcdef`
	b := s[0:2]
	c := s[2:5]
	fmt.Println(b,"---",c)
}

func Test_Buf(t *testing.T) {
	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use big endian: \n", testInt)
	testBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)

}