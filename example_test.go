package wangdiantong_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/treelab/wangdiantong-go-sdk"
)

var server *httptest.Server

func setup() {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"code":        0,
			"message":     "ok",
			"total_count": "210",
			"goods_list": map[string]string{
				"rec_id":            "17",
				"shop_no":           "manmao2test",
				"shop_name":         "manmao2test",
				"platform_id":       "127",
				"match_target_id":   "4",
				"match_target_type": "1",
				"api_goods_id":      "0000",
				"api_spec_id":       "0000",
				"api_goods_name":    "测试用例0",
				"api_spec_name":     "5474",
				"modified":          "2019-01-03 17:29:34",
				"api_spec_no":       "",
				"outer_id":          "11",
				"spec_outer_id":     "qqq",
				"stock_num":         "496.0000",
				"price":             "20.0000",
				"cid":               "",
				"pic_url":           "",
				"is_deleted":        "2",
				"hold_stock":        "0.0000",
				"hold_stock_type":   "0",
				"is_auto_listing":   "1",
				"is_auto_delisting": "1",
				"status":            "1",
				"merchant_no":       "qqq",
				"merchant_name":     "5474",
				"merchant_code":     "dfd",
			},
		})
	}))
}

func tearDown() {
	server.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func Exampleclient_Execute() {
	c, err := wangdiantong.New(
		&wangdiantong.Options{
			AppKey:    "test",
			AppSecret: "test",
			SID:       "test",
			BaseURL:   server.URL,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	params := map[string]string{
		"start_time": "2019-01-01 00:00:00",
		"end_time":   "2019-01-30 00:00:00",
		"page_no":    "0",
		"page_size":  "10",
	}

	resp, err := c.Execute("/vip_api_goods_query.php", params)

	if err != nil {
		log.Fatalln(err)
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("code:", res["code"])
	fmt.Println("message:", res["message"])
	fmt.Println("total_count:", res["total_count"])
	fmt.Println("rec_id:", res["goods_list"].(map[string]interface{})["rec_id"])
	fmt.Println("shop_no:", res["goods_list"].(map[string]interface{})["shop_no"])
	fmt.Println("shop_name:", res["goods_list"].(map[string]interface{})["shop_name"])
	// Output:
	// code: 0
	// message: ok
	// total_count: 210
	// rec_id: 17
	// shop_no: manmao2test
	// shop_name: manmao2test
}
