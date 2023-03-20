package http

import (
	"testing"
	"time"
)

func TestHttpGet(t *testing.T) {
	HttpGet("http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=sh000001&scale=60&datalen=1000000")
}

func TestHttpHead(t *testing.T) {
	url := "https://finance.sina.com.cn/realstock/company/klc_td_sh.txt"
	ts := "Wed, 21 Dec 2022 09:59:52 GMT"
	lastModified, _ := time.Parse(time.RFC1123, ts)
	header := map[string]any{
		IfModifiedSince: lastModified,
	}
	Request(url, "get", header)
}
