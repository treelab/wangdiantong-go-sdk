package wangdiantong

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// see: https://open.wangdian.cn/open/guide?path=guide_signsf
func TestSignRequest(t *testing.T) {
	appSecret := "12345"

	params := map[string]string{
		"appkey":     "test2-xx",
		"page_no":    "0",
		"end_time":   "2016-08-01 13:00:00",
		"start_time": "2016-08-01 12:00:00",
		"page_size":  "40",
		"sid":        "test2",
		"timestamp":  "1470042310",
	}

	expected := "ad4e6fe037ea6e3ba4768317be9d1309"
	if out := signRequest(params, appSecret); out != expected {
		t.Errorf("\ngot  %v\nwant %v", out, expected)
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		input    map[string]string
		expected string
	}{
		{map[string]string{"foo": "bar"}, "foo=bar"},
		// encode() will sort params using key in ascending order
		{map[string]string{"foo": "bar", "bar": "baz"}, "bar=baz&foo=bar"},
		{map[string]string{"a": "a?b!c", "bar": "baz"}, "a=a%3Fb%21c&bar=baz"},
	}

	for _, test := range tests {
		output := encode(test.input)
		if output != test.expected {
			t.Errorf("\ngot: %v\nwant %v", output, test.expected)
		}
	}
}

func TestPost(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctype := req.Header.Get("Content-Type")
		if ctype != "application/x-www-form-urlencoded" {
			t.Errorf("Content-Type want: %v, got %v", "application/x-www-form-urlencoded", ctype)
		}

		err := req.ParseForm()
		if err != nil {
			t.Fatal(err)
		}

		if req.PostForm["appkey"][0] != "appkey_value" {
			t.Errorf("appKey want %v, got %v", "appkey_value", req.PostForm["appkey"])
		}

		if req.PostForm["sid"][0] != "sid_value" {
			t.Errorf("appKey want %v, got %v", "sid_value", req.PostForm["sid"])
		}

		if req.PostForm["timestamp"][0] != "timestamp_value" {
			t.Errorf("appKey want %v, got %v", "timestamp_value", req.PostForm["timestamp"])
		}

		if req.PostForm["sign"][0] != "sign_value" {
			t.Errorf("appKey want %v, got %v", "sign_value", req.PostForm["sign"])
		}

		if req.Method != http.MethodPost {
			t.Errorf("httpMethod want: %v, got %v", http.MethodPost, req.Method)
		}

		fmt.Fprintf(w, "ok")
	}))

	defer srv.Close()
	c, err := New(&Options{
		AppKey:    "1",
		AppSecret: "1",
		SID:       "1",
		BaseURL:   "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.post(srv.URL, map[string]string{
		"appkey":    "appkey_value",
		"sid":       "sid_value",
		"timestamp": "timestamp_value",
		"sign":      "sign_value",
	})

	if err != nil {
		t.Fatal(err)
	}
}
