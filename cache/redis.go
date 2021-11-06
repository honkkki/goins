package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var r redis.Conn

func InitRedis() (err error) {
	host := "127.0.0.1"
	port := "6379"

	if r == nil {
		once := sync.Once{}
		once.Do(func() {
			r, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		})
	}

	return err
}

func GetCookie() (ret string, err error) {
	err = InitRedis()
	if err != nil {
		return "", err
	}

	ret, err = redis.String(r.Do("get", "ins-cookie"))
	return
}
