package wangdiantong

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"time"
)

type client struct {
	AppKey    string
	AppSecret string
	SID       string
	BaseURL   string
}

func New(
	appKey string,
	appSecret string,
	SID string,
	baseURL string,
) *client {
	return &client{
		AppKey:    appKey,
		AppSecret: appSecret,
		SID:       SID,
		BaseURL:   baseURL,
	}
}

func (c *client) Execute(relativeURL string, params map[string]string) string {
	params["appkey"] = c.AppKey
	params["sid"] = c.SID
	params["timestamp"] = fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))
	params["sign"] = c.signRequest(params)

	// TODO: make HTTP Post request
	return ""
}

func (c *client) signRequest(params map[string]string) string {
	// 第一步：对所有请求参数按照键名进行正序排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 第二步：把所有参数名和参数值串在一起,
	var bf bytes.Buffer
	for _, key := range keys {
		if key == "sign" {
			continue
		}

		if bf.Len() > 0 {
			bf.WriteString(";")
		}

		// example:
		// input: {"appkey" : "test2-xx"}
		// output: "06-appkey:0008-test2-xx"
		str := fmt.Sprintf("%02d-%s:%04d-%s", len(key), key, len(params[key]), params[key])
		bf.WriteString(str)
	}

	bf.WriteString(c.AppSecret)
	fmt.Println(bf.String())

	// 第三步：使用MD5哈希
	hash := md5.Sum(bf.Bytes())

	// 第四步：把二进制转化为大写的十六进制
	hexStr := hex.EncodeToString(hash[:])

	return hexStr
}
