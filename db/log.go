package db

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"log"
	"time"
)

type Log struct {
	LogLevel logger.LogLevel
}

func (l *Log) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *Log) Info(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("Info implement me")
}

func (l *Log) Warn(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("Warn implement me")
}

func (l *Log) Error(ctx context.Context, s string, i ...interface{}) {
	//TODO implement me
	panic("Error implement me")
}

func (l *Log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	log.Println(fmt.Sprintf("sql:%s,rows:%d", sql, rows))
	if err != nil {
		utils.FileWithLineNum()
		field := zap.Error(err)
		log.Println(field.Key, field.Interface)

	}
	return
}

func NewLog() logger.Interface {
	return new(Log)
}
