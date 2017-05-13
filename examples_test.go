package ht

import (
	. "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get(t *testing.T) {
	expected := "Hello World"
	server := StringServer(expected, 200)
	defer server.Close()

	result, err := Expect200(Get(server.URL)).String()
	NoError(t, err)
	Equal(t, expected, result)

	_, err = Expect500(Get(server.URL)).String()
	Error(t, err)
}

func Test_Get500(t *testing.T) {
	expected := "Hello World"
	server := StringServer(expected, 500)
	defer server.Close()

	_, err := Expect200(Get(server.URL)).String()
	Error(t, err)

	_, err = Expect2xx(Get(server.URL)).String()
	Error(t, err)

	_, err = Expect404(Get(server.URL)).String()
	Error(t, err)

	result, err := Expect500(Get(server.URL)).String()
	NoError(t, err)
	Equal(t, expected, result)
}

func Test_GetJson(t *testing.T) {
	expected := `{"some": "value"}`
	server := StringServer(expected, 200, TypeJson)
	defer server.Close()

	result := map[string]string{}
	err := Expect200(Get(server.URL)).Json(&result)
	NoError(t, err)
	Equal(t, "value", result["some"])
}

func StringServer(responseBody string, code int, header ...string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetHeader(w.Header(), header...)
		w.WriteHeader(code)
		w.Write([]byte(responseBody))
	}))
}
