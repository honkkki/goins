package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/honkkki/goins/config"
	"log"
)

var r redis.Conn

func init() {
	res := config.GetRedisConfig()
	host := res["host"]
	port := res["port"]

	var err error
	r, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("redis init failed: %v", err)
	}
}

func GetCookie() (ret string, err error) {
	ret, err = redis.String(r.Do("get", "ins-cookie"))
	return
}

func GetTag() (ret string, err error) {
	ret, err = redis.String(r.Do("get", "tag"))
	return
}
