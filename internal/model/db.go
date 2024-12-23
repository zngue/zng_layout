package model

import (
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zngue/zng_app/db"
	"github.com/zngue/zng_app/db/mysql"
	"github.com/zngue/zng_app/db/redis"
	"github.com/zngue/zng_layout/internal/conf"
	"github.com/zngue/zng_layout/pkg/util"
	"gorm.io/gorm"
	"os"
)

func NewDB(bootstrap *conf.Bootstrap) (conn *gorm.DB, err error) {
	var config = bootstrap.Mysql
	database := os.Getenv("DB_DATABASE")
	if database != "" {
		config.Database = database
	}
	conn, err = db.NewDB(
		mysql.DataWithDatabase(config.Database),
		mysql.DataWithHost(config.Host),
		mysql.DataWithPassword(config.Password),
		mysql.DataWithPort(int(config.Port)),
		mysql.DataWithUserName(config.Username),
		mysql.DataWithLogger(true),
		mysql.DataWithLoggerConfig(util.LogConfig()),
	)
	if err != nil {
		return
	}
	level := os.Getenv("LOG_LEVEL")
	if level == "debug" {
		conn = conn.Debug()
	}
	return

}

func NewRedis(bootstrap *conf.Bootstrap) (*redis2.Client, func(), error) {
	var config = bootstrap.Redis
	return db.NewRedis(
		redis.DataWithHost(config.Host),
		redis.DataWithPort(int(config.Port)),
		redis.DataWithPassword(config.Password),
	)
}
