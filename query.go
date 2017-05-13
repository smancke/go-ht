package ht

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(urlStr string, header ...string) (*http.Response, error) {
	return Do("GET", urlStr, nil, header...)
}

func Delete(urlStr string, header ...string) (*http.Response, error) {
	return Do("DELETE", urlStr, nil, header...)
}

func Post(urlStr string, body string, header ...string) (*http.Response, error) {
	if !containsContentType(header) {
		contentType := http.DetectContentType([]byte(body))
		header = append(header, ContentType(contentType))
	}

	return Do("POST", urlStr, bytes.NewBufferString(body), header...)
}

func PostJson(urlStr string, body interface{}, header ...string) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	if !containsContentType(header) {
		header = append(header, ContentType("application/json"))
	}

	return Do("POST", urlStr, bytes.NewBuffer(bodyBytes), header...)
}

func PostForm(urlStr string, body url.Values, header ...string) (*http.Response, error) {
	if !containsContentType(header) {
		header = append(header, ContentType("application/x-www-form-urlencoded"))
	}

	return Do("POST", urlStr, bytes.NewBufferString(body.Encode()), header...)
}

func containsContentType(header []string) bool {
	for _, h := range header {
		if strings.HasPrefix(h, "Content-Type:") {
			return true
		}
	}
	return false
}

func Do(method string, urlStr string, body io.Reader, header ...string) (*http.Response, error) {
	r, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	err = SetHeader(r.Header, header...)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(r)
}
