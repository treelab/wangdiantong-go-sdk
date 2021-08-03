## 旺店通 Golang SDK

This library is a Go implementation for [旺店通 SDK](https://open.wangdian.cn/qyb/open/abut/sdk_download)

## Install

```bash
$ go get github.com/treelab/wangdiantong-go-sdk
```

## Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/treelab/wangdiantong-go-sdk"
)

func main() {
	wdClient, err := wangdiantong.New(&wangdiantong.Options{
		AppKey:     "YOUR_APP_KEY",
		AppSecret:  "YOUR_APP_SECRET",
		SID:        "YOUR_SID",
		BaseURL:    "https://sandbox.wangdian.cn/openapi2",
		HttpClient: &http.Client{},
	})

	if err != nil {
		log.Fatal(err)
	}

	params := map[string]string{
		"start_time": "2019-01-01 00:00:00",
		"end_time":   "2019-01-30 00:00:00",
		"page_no":    "0",
		"page_size":  "10",
	}

	resp, err := wdClient.Execute("/vip_api_goods_query.php", params)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}

```

## License

Copyright (c) 2021 Treelab, released under MIT License
