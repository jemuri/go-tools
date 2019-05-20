package log

import (
	"context"
	"os"

	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

//ExistObj ExistObj
//type ExistObj int

const (
	//EntryInstance EntryInstance
	EntryInstance = "EntryInstance"
	//fileKey file Key
	fileKey = "file"
	//Project env key
	Project = "PROJECT"
	//FuncNameKey FuncNameKey
	FuncNameKey = "func_name"
	// ProjectKey ProjectKey
	ProjectKey = "project"
	// DefaultEnv DefaultEnv
	DefaultEnv = "aegis-c2c"
)

//LoggerInit 初始化日志记录器
func LoggerInit(ctx context.Context, args map[string]interface{}) context.Context {
	if ctx == nil {
		panic("No context")
	}
	entry, isOK := ctx.Value(EntryInstance).(*logrus.Entry) //从context中获取log
	if !isOK {
		entry = logrus.NewEntry(logrus.New()) //取不到时创建一个新的entry日志记录器
	}
	args[FuncNameKey] = getFuncName()                //获取gRPC接口名称
	args[ProjectKey] = getEnvProject()               //项目应用部署名称,再也不用担心重新部署后找不到日志了
	entry = entry.WithFields(logrus.Fields(args))    //将业务线标示信息(eg.apply_id)埋入该日志记录器中
	entry.Logger.Formatter = &logrus.JSONFormatter{} //日志实例处 设置  才生效

	return context.WithValue(ctx, EntryInstance, entry)
}

func getEnvProject() string {
	project := os.Getenv(Project)
	if project == "" {
		project = DefaultEnv
	}

	return project
}

//AddFields 新增业务埋点
func AddFields(ctx context.Context, args map[string]interface{}) context.Context {
	entry, isOK := ctx.Value(EntryInstance).(*logrus.Entry) //从context中获取log
	if !isOK {
		entry = logrus.NewEntry(logrus.New()) //取不到时创建一个新的entry日志记录器
		entry.Logger.Formatter = &logrus.JSONFormatter{}
	}
	entry = entry.WithFields(logrus.Fields(args)) //将业务线标示信息(eg.apply_id)埋入该日志记录器中

	return context.WithValue(ctx, EntryInstance, entry)
}

//getLineNumber get line number
func getLineNumber() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%d", file, line)
	}

	return ""
}

//getFuncName get function name
func getFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		names := strings.Split(funcName, ".")
		if len(names) > 0 {
			return names[len(names)-1]
		}
	}

	return ""
}

//Debug 测试日志记录
func Debug(ctx context.Context, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Debug(args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Debug(args...)
}

//Debugf 测试日志记录
func Debugf(ctx context.Context, format string, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Debugf(format, args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Debugf(format, args...)
}

//Info 日志记录
func Info(ctx context.Context, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Info(args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Info(args...)
}

//Infof 日志记录
func Infof(ctx context.Context, format string, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Infof(format, args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Infof(format, args...)
}

//Warn 警告日志记录
func Warn(ctx context.Context, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Warn(args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Warn(args...)
}

//Warnf 警告日志记录
func Warnf(ctx context.Context, format string, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Warnf(format, args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Warnf(format, args...)
}

//Error 错误日志记录
func Error(ctx context.Context, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Error(args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Error(args...)
}

//Errorf 错误日志记录
func Errorf(ctx context.Context, format string, args ...interface{}) {

	entry := ctx.Value(EntryInstance)
	if entry == nil {
		entryNew := logrus.NewEntry(logrus.New())
		entryNew.Warn("can't find entry")
		entryNew.Errorf(format, args...)
		return
	}
	entry.(*logrus.Entry).WithFields(map[string]interface{}{fileKey: getLineNumber()}).Errorf(format, args...)
}
