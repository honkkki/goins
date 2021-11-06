package config

import (
	"fmt"
	"github.com/honkkki/goins/utils"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var rootDir string

func init()  {
	GetRoot()
}

func GetRoot() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var recur func(wd string)
	recur = func(wd string) {
		if utils.Exist(wd + "/utils") {
			rootDir = wd
			return
		}

		recur(filepath.Dir(wd))
	}

	recur(wd)
}

func GetSavePath() string {
	fmt.Println(rootDir)
	viper.SetConfigFile(rootDir + "/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.WatchConfig()
	return viper.GetString("file.path")
}
