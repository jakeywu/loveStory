package cache

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"strings"
)

var c cache.Cache

func InitRedis() {
	var err error
	var config string
	conn := strings.Join([]string{beego.AppConfig.String("redisCache::Host"), beego.AppConfig.String("redisCache::Port")}, ":")
	password := beego.AppConfig.String("redisCache::Pwd")
	if password != "" {
		config = fmt.Sprintf(`{"key":"beeApi","conn":"%s", "dbNum": "%s", "password": "%s"}`, conn, beego.AppConfig.String("redisCache::DbNum"), password)
	} else {
		config = fmt.Sprintf(`{"key":"beeApi","conn":"%s", "dbNum": "%s"}`, conn, beego.AppConfig.String("redisCache::DbNum"))
	}
	c, err = cache.NewCache("redis", config)
	if err != nil {
		panic("cache init default, please check redisCache config")
	}
}
