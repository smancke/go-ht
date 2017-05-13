package ht

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	TypeJson  = "Content-Type: application/json"
	TypePlain = "Content-Type: text/plain"
	TypeHtml  = "Content-Type: text/html"
)

func SetHeader(rh http.Header, header ...string) error {
	for _, h := range header {
		pairs := strings.SplitN(h, ": ", 1)
		if len(pairs) == 2 {
			rh.Set(pairs[0], pairs[1])
		} else {
			return fmt.Errorf("invalid header format %q", h)
		}
	}
	return nil
}

func ContentType(ct string) string {
	return "Content-Type: " + ct
}

func Bearer(secret string) string {
	return "Authorization: Bearer " + secret
}
