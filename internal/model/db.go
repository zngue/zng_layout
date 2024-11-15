package model

import (
	redis2 "github.com/redis/go-redis/v9"
	"github.com/zngue/zng_app/db"
	"github.com/zngue/zng_app/db/mysql"
	"github.com/zngue/zng_app/db/redis"
	"github.com/zngue/zng_layout/internal/conf"
	"gorm.io/gorm"
)

func NewDB(bootstrap *conf.Bootstrap) (*gorm.DB, error) {
	var config = bootstrap.Mysql
	return db.NewDB(
		mysql.DataWithDatabase(config.Database),
		mysql.DataWithHost(config.Host),
		mysql.DataWithPassword(config.Password),
		mysql.DataWithPort(int(config.Port)),
		mysql.DataWithUserName(config.Username),
	)

}

func NewRedis(bootstrap *conf.Bootstrap) (*redis2.Client, func(), error) {
	var config = bootstrap.Redis
	return db.NewRedis(
		redis.DataWithHost(config.Host),
		redis.DataWithPort(int(config.Port)),
		redis.DataWithPassword(config.Password),
	)
}
