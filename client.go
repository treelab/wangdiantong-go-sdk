package wangdiantong

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Options struct {
	AppKey     string
	AppSecret  string
	SID        string
	BaseURL    string
	HttpClient *http.Client
}

func (opt *Options) init() error {
	if opt == nil {
		return errors.New("wangdiantong: opt cannot be nil")
	}

	if opt.AppKey == "" {
		return fmt.Errorf("wangdiantong: invalid AppKey: %q", opt.AppKey)
	}

	if opt.AppSecret == "" {
		return fmt.Errorf("wangdiantong: invalid AppSecret: %q", opt.AppKey)
	}

	if opt.SID == "" {
		return fmt.Errorf("wangdiantong: invalid SID: %q", opt.AppKey)
	}

	if opt.BaseURL == "" {
		return fmt.Errorf("wangdiantong: invalid BaseUrl: %q", opt.AppKey)
	}

	if opt.HttpClient == nil {
		opt.HttpClient = &http.Client{}
	}

	return nil
}

type client struct {
	appKey     string
	appSecret  string
	sid        string
	baseURL    string
	httpClient *http.Client
}

func New(opt *Options) (*client, error) {
	if err := opt.init(); err != nil {
		return nil, err
	}

	return &client{
		appKey:     opt.AppKey,
		appSecret:  opt.AppSecret,
		sid:        opt.SID,
		baseURL:    opt.BaseURL,
		httpClient: opt.HttpClient,
	}, nil
}

func (c *client) Execute(relativeURL string, params map[string]string) (*http.Response, error) {
	params["appkey"] = c.appKey
	params["sid"] = c.sid
	params["timestamp"] = fmt.Sprint(time.Now().Unix())
	params["sign"] = signRequest(params, c.appKey)

	return c.post(c.baseURL+relativeURL, params)
}

func signRequest(params map[string]string, appSecret string) string {
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

	bf.WriteString(appSecret)

	// 第三步：使用MD5哈希
	hash := md5.Sum(bf.Bytes())

	// 第四步：把二进制转化为大写的十六进制
	hexStr := hex.EncodeToString(hash[:])

	return hexStr
}

func (c *client) post(url string, data map[string]string) (*http.Response, error) {
	encoded := encode(data)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(encoded))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.httpClient.Do(req)
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func encode(params map[string]string) string {
	var buf strings.Builder
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := params[k]
		keyEscaped := url.QueryEscape(k)
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(v))
	}

	return buf.String()
}
