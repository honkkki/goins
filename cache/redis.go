package cache

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/honkkki/goins/config"
	"log"
	"sync"
)

var r redis.Conn

func init() {
	res := config.GetRedisConfig()
	host := res["host"]
	port := res["port"]

	if r == nil {
		once := sync.Once{}
		once.Do(func() {
			var err error
			r, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
			if err != nil {
				log.Printf("redis init failed: %v", err)
				return
			}
		})
	}
}

func GetCookie() (ret string, err error) {
	if r == nil {
		return "", errors.New("redis init failed")
	}
	
	ret, err = redis.String(r.Do("get", "ins-cookie"))
	return
}
