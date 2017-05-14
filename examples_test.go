package ht

import (
	"github.com/smancke/go-ht/htest"
	. "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Fetch(t *testing.T) {
	server := htest.StringServer("Hello World", 200)
	defer server.Close()

	result, err := Fetch(server.URL)
	NoError(t, err)
	Equal(t, "Hello World", result)

}

func Test_Fetch_Error(t *testing.T) {
	server := htest.StringServer("Hello World", 500)
	defer server.Close()

	_, err := Fetch(server.URL)
	Error(t, err)
}

func Test_FetchJson(t *testing.T) {
	server := htest.StringServer(`{"some": "value"}`, 200, TypeJson)
	defer server.Close()

	result := map[string]string{}
	err := FetchJson(server.URL, &result)
	NoError(t, err)
	Equal(t, "value", result["some"])

}

func Test_FetchJson_Error(t *testing.T) {
	server := htest.StringServer(`{"some": "value"`, 200, TypeJson)
	defer server.Close()

	result := map[string]string{}
	err := FetchJson(server.URL, &result)
	Error(t, err)
}

func Test_Get(t *testing.T) {
	server := htest.StringServer("Hello World", 200)
	defer server.Close()

	result, err := Expect200(Get(server.URL)).String()
	NoError(t, err)
	Equal(t, "Hello World", result)

	_, err = Expect500(Get(server.URL)).String()
	Error(t, err)
}

func Test_Get500(t *testing.T) {
	expected := "Hello World"
	server := htest.StringServer(expected, 500)
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
	server := htest.StringServer(`{"some": "value"}`, 200, TypeJson)
	defer server.Close()

	result := map[string]string{}
	err := Expect200(Get(server.URL)).Json(&result)
	NoError(t, err)
	Equal(t, "value", result["some"])
}
