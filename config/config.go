package config

import (
	"github.com/honkkki/goins/utils"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var (
	rootDir string
	vip *viper.Viper
)

func init()  {
	GetRoot()
	if vip == nil {
		vip = viper.New()
		vip.SetConfigFile(rootDir + "/config.yaml")
	}

	err := vip.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	vip.WatchConfig()
}

func GetRoot() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var recur func(wd string)
	recur = func(wd string) {
		if utils.Exist(wd + "/config.yaml") {
			rootDir = wd
			return
		}

		if wd == filepath.Dir(wd) {
			log.Fatal("config.yaml not found.")
		}

		recur(filepath.Dir(wd))
	}

	recur(wd)
}

func GetSavePath() string {
	return vip.GetString("file.path")
}

func GetRedisConfig() map[string]string{
	m := make(map[string]string)
	m["host"] = vip.GetString("redis.host")
	m["port"] = vip.GetString("redis.port")
	return m
}
