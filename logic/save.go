package logic

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/honkkki/goins/config"
	"github.com/honkkki/goins/utils"
)

func Save(resource map[string]string, username string) {
	log.Println(resource)
	baseDir := config.GetSavePath()
	dir := filepath.Join(baseDir, username)

	if !utils.Exist(dir) {
		// create dir
		os.Mkdir(dir, 0777)
	}

	var wg sync.WaitGroup
	count := len(resource)
	wg.Add(count)
	fmt.Println(username)
	for k, v := range resource {
		go func(id string, url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			defer res.Body.Close()
			stream, _ := ioutil.ReadAll(res.Body)
			fileName := id + ".jpg"
			path := filepath.Join(baseDir, username, fileName)
			ioutil.WriteFile(path, stream, 0644)
			fmt.Println("download item finish!")
		}(k, v)
	}

	wg.Wait()
	fmt.Println("all finish!")
}
