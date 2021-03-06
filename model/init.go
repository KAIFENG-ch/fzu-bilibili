package model

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Database(dsn string) {
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)
	DB = db
	err = DB.Set("gorm:table_option", "ENGINE=InnoDB").
		AutoMigrate(&User{}, &Video{}, &Comments{}, &Replies{}, &Barrage{})
	if err != nil {
		panic(err)
	}
}

var RDB *redis.Client

func RedisDb(addr string, password string, db int) (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})
	RDB = rdb
	_,err = RDB.Ping().Result()
	if err != nil{
		return err
	}
	return nil
}
