package wangdiantong

import "testing"

// see: https://open.wangdian.cn/open/guide?path=guide_signsf
func TestSignRequest(t *testing.T) {
	client := New("", "12345", "", "")
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
	if out := client.signRequest(params); out != expected {
		t.Errorf("\ngot  %v\nwant %v", out, expected)
	}
}
