package http

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"github.com/mymmsc/gox/api"
	"github.com/mymmsc/gox/logger"
	"io"
	"net/http"
	URL "net/url"
	"strings"
	"time"
)

const (
	ContentEncoding = "Content-Encoding"
	LastModified    = "Last-Modified"
	IfModifiedSince = "If-Modified-Since"
)

var (
	TimeZero = time.Unix(0, 0)
)

type Response struct {
	StatusCode    int
	ContentLength int
	LastModified  time.Time
	Body          []byte
}

func HttpRequest(url string, method string) ([]byte, error) {
	data, lastModified, err := Request(url, method)
	_ = lastModified
	return data, err
}

func HttpGet(url string) ([]byte, error) {
	return HttpRequest(url, "get")
}

// Request http request, 支持传入header
func Request(url string, method string, header ...map[string]any) (data []byte, lastModified time.Time, err error) {
	u, err := URL.Parse(url)
	if err != nil {
		return nil, TimeZero, err
	}
	reqHeader := make(map[string]string)
	reqHeader["Accept"] = "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	reqHeader["Accept-Encoding"] = "gzip, deflate"
	reqHeader["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	reqHeader["Cache-Control"] = "no-cache"
	reqHeader["Connection"] = "keep-alive"
	reqHeader["Host"] = u.Host
	reqHeader["Pragma"] = "no-cache"
	reqHeader["Upgrade-Insecure-Requests"] = "1"
	reqHeader["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	if len(header) > 0 {
		mapHeader := header[0]
		for k, v := range mapHeader {
			switch val := v.(type) {
			case time.Time:
				val = val.UTC()
				reqHeader[k] = val.Format(time.RFC1123)
			}
		}
	}

	client := &http.Client{}
	request, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		return nil, TimeZero, err
	}
	for key, v := range reqHeader {
		request.Header.Add(key, v)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, TimeZero, err
	}
	lm := response.Header.Get(LastModified)
	if response.StatusCode == http.StatusNotModified && !api.IsEmpty(lm) {
		return nil, TimeZero, nil
	}
	lastModified, err = time.Parse(time.RFC1123, lm)
	defer api.CloseQuietly(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, TimeZero, err
	}
	contentEncoding := response.Header.Get(ContentEncoding)
	var reader io.ReadCloser = nil
	if len(contentEncoding) > 0 {
		contentEncoding = strings.ToLower(contentEncoding)
		switch contentEncoding {
		case "gzip":
			reader, err = gzip.NewReader(bytes.NewBuffer(body))
			if err != nil {
				logger.Error(err)
				reader = nil
			}
		case "deflate":
			reader = flate.NewReader(bytes.NewReader(body))
		}
	}
	if reader != nil {
		defer api.CloseQuietly(reader)
		body, err = io.ReadAll(reader)
		if err != nil {
			return nil, TimeZero, err
		}
	}
	return body, lastModified, nil
}
