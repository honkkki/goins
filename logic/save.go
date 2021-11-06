package logic

import (
	"fmt"
	"github.com/honkkki/goins/config"
	"github.com/honkkki/goins/utils"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func Save(resource map[string]string, username string) {
	baseDir := config.GetSavePath()
	dir := filepath.Join(baseDir, username)

	if !utils.Exist(dir) {
		// create dir
		os.Mkdir(dir, 0644)
	}

	var wg sync.WaitGroup
	count := len(resource)
	wg.Add(count)
	fmt.Println(username)
	for k, v := range resource {
		go func(id string, url string) {
			defer wg.Done()
			res, _ := http.Get(url)
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
