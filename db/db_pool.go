package db

import (
	"errors"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)



var (
	connCache sync.Map
	lock      sync.Mutex
)

type Config struct {
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifeTime time.Duration
}

type dbInfo struct {
	db  *gorm.DB
	cfg *Config
}

type GormLogger interface {
	Print(v ...interface{})
	DoExplain(time.Duration, *gorm.Scope)
}

var (
	// 默认参数：最大连接数64，最大空闲连接：16，连接最大存活时间：10分钟
	defaultCfg = Config{MaxIdleConn: 16, MaxOpenConn: 64, ConnMaxLifeTime: 10 * time.Minute}
	GormLog    GormLogger
)

// InitDB _
func InitDB(name, dsn string) error {
	return CreateDBWithConfig(name, dsn, defaultCfg)
}

// CreateDBWithConfig _
func CreateDBWithConfig(name, dsn string, cfg Config) error {

	lock.Lock()
	defer lock.Unlock()

	if name == "" {
		return errors.New("name is empty")
	}
	if dsn == "" {
		return errors.New("dsn is empty")
	}
	//if GormLog == nil {
	//	return errors.New("should call SetLogger first")
	//}

	var (
		conn *gorm.DB
		err  error
	)

	conn, err = gorm.Open("mysql", dsn)
	if err != nil {
		return errors.New("failed to connect database")
	}

	conn.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	conn.DB().SetMaxIdleConns(cfg.MaxOpenConn)
	conn.DB().SetConnMaxLifetime(cfg.ConnMaxLifeTime)

	// 设置logrus log
	//conn.SetLogger(GormLog)

	// 初始化查询开始时间，用于统计查询耗时
	//conn.Callback().Query().Before("gorm:query").Register("initQueryTime", CallbackInitTime)
	// 执行explain 语句
	//conn.Callback().Query().Register("explain", CallbackExplain)

	connCache.Store(name, &dbInfo{conn, &cfg})
	return nil
}

// GetDB _
func GetDB(name string) (*gorm.DB, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	temp, ok := connCache.Load(name)
	if !ok {
		return nil, errors.New("should call createDB first")
	}
	val := temp.(*dbInfo)
	drain(val)

	return val.db, nil
}

func drain(val *dbInfo) {

	var (
		err error
	)

	for i := 0; i < val.cfg.MaxIdleConn; i++ {
		if err = val.db.DB().Ping(); err == nil {
			return
		}
	}
	return
}

// SetLogger 设置gorm log
func SetLogger(gormLog GormLogger) {
	GormLog = gormLog
}


// CallbackInitTime 初始化查询时间
func CallbackInitTime(scope *gorm.Scope) {
	scope.Set("start_time", time.Now())
}

// CallbackExplain _
func CallbackExplain(scope *gorm.Scope) {

	// 获取sql查询开始时间
	temp, ok := scope.Get("start_time")
	if !ok {
		return
	}
	start, ok := temp.(time.Time)
	if !ok {
		return
	}

	duration := time.Since(start)
	GormLog.DoExplain(duration, scope)
}

// deprecated IsEmptyDB 如果不使用事务管理,不需要传递db参数
func IsEmptyDB(db ...*gorm.DB) (*gorm.DB,error) {
	if db == nil && len(db)==0{
		dbNew, err := GetDB("master")
		if err != nil {
			return nil,err
		}
		return dbNew,nil
	}
	return db[0],nil
}
