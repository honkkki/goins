package main

import (
	"encoding/json"
	"fmt"
	"github.com/honkkki/goins/cache"
	"github.com/honkkki/goins/insmodel"
	"github.com/honkkki/goins/logic"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	for {
		// get url from cmd
		fmt.Println("type url:")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Println("get url failed:", err)
			continue
		}

		if input == "q" {
			fmt.Println("see you.")
			break
		}

		tag, err := cache.GetTag()
		if err != nil {
			fmt.Println("cant get tag:", err)
			return
		}

		url := input + tag
		fmt.Println("fetching url...")
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)

		// get cookie from redis
		cookie, err := cache.GetCookie()
		if err != nil {
			fmt.Println("cant get ins cookie:", err)
			return
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")
		req.Header.Set("Cookie", cookie)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if code := resp.StatusCode; code != 200 {
			fmt.Println("response http code error:", code)
			return
		}

		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		ig := insmodel.IG{}
		err = json.Unmarshal(data, &ig)
		if err != nil {
			fmt.Println("json parse failed:", err)
			return
		}

		imgMap := make(map[string]string)
		imgUsername := ig.Graphql.ShortcodeMedia.Owner.Username
		resType := ig.Graphql.ShortcodeMedia.Typename

		switch resType {
		case "GraphImage":
			imgMap[ig.Graphql.ShortcodeMedia.ID] = ig.Graphql.ShortcodeMedia.DisplayURL
		case "GraphSidecar":
			imgResource := ig.Graphql.ShortcodeMedia.EdgeSidecarToChildren.Edges
			for _, v := range imgResource {
				imgMap[v.Node.ID] = v.Node.DisplayURL
			}
		default:
			log.Println("error resource type.")
			continue
		}

		logic.Save(imgMap, imgUsername)
	}
}
