package log

import (
	"context"
	"reflect"
	"testing"
)

func TestLoggerInit(t *testing.T) {
	type args struct {
		ctx  context.Context
		args map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "case1",
			args: args{
				ctx:  context.Background(),
				args: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoggerInit(tt.args.ctx, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//type logFileWriter struct {
//	file *os.File
//	//write count
//	size int64
//}
//
//var wg sync.WaitGroup
//
//func main() {
//	//log.SetFormatter(&log.JSONFormatter{})
//	file, err := os.OpenFile(fmt.Sprintf("logs/ying_%s.log", time.Now().Format("2006_01_02")), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
//	if err != nil {
//		log.Fatal("log  init failed")
//	}
//
//	info, err := file.Stat()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fileWriter := logFileWriter{file, info.Size()}
//	log.SetOutput(&fileWriter)
//	log.Info("start.....")
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go logTest(i)
//	}
//	log.Warn("waitting...")
//	wg.Wait()
//}
//func (p *logFileWriter) Write(data []byte) (n int, err error) {
//	if p == nil {
//		return 0, errors.New("logFileWriter is nil")
//	}
//	if p.file == nil {
//		return 0, errors.New("file not opened")
//	}
//	n, e := p.file.Write(data)
//	p.size += int64(n)
//	//文件最大 64K byte
//	if p.size > 1024*64 {
//		p.file.Close()
//		fmt.Println("log file full")
//		p.file, _ = os.OpenFile("./mylog"+strconv.FormatInt(time.Now().Unix(), 10), os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0600)
//		p.size = 0
//	}
//	return n, e
//}
//
//func logTest(id int) {
//	for i := 0; i < 100; i++ {
//		log.Info("Thread:", id, " value:", i)
//		time.Sleep(10 * time.Millisecond)
//	}
//	wg.Done()
//}

func TestLogInit(t *testing.T) {
	ctx := LogInit(context.Background(), map[string]interface{}{})
	for i := 0; i < 100; i++ {
		Infof(ctx, "[wing-TestLogInit]第: %d条", i)
	}
}
