package ht

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Res struct {
	res *http.Response
	err error
}

func NewRes(httpRes *http.Response, err error) *Res {
	return &Res{httpRes, err}
}

func Expect200(httpRes *http.Response, err error) *Res {
	return NewRes(httpRes, err).ExpectCode(200)
}

func Expect201(httpRes *http.Response, err error) *Res {
	return NewRes(httpRes, err).ExpectCode(201)
}

func Expect404(httpRes *http.Response, err error) *Res {
	return NewRes(httpRes, err).ExpectCode(404)
}

func Expect500(httpRes *http.Response, err error) *Res {
	return NewRes(httpRes, err).ExpectCode(500)
}

func Expect2xx(httpRes *http.Response, err error) *Res {
	return NewRes(httpRes, err).Expect2xx()
}

func (r *Res) Expect2xx() *Res {
	if r.err == nil {
		if r.res.StatusCode < 200 || r.res.StatusCode > 299 {
			r.err = fmt.Errorf("expect http code 2xx, but got %v", r.res.StatusCode)
		}
	}
	return r
}

func (r *Res) ExpectCode(code int) *Res {
	if r.err == nil {
		if r.res.StatusCode < code || r.res.StatusCode > code {
			r.err = fmt.Errorf("expect http code %v, but got %v", code, r.res.StatusCode)
		}
	}
	return r
}

func (r *Res) Json(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	defer r.res.Body.Close()
	dec := json.NewDecoder(r.res.Body)
	return dec.Decode(v)
}

func (r *Res) String() (string, error) {
	b, err := r.Bytes()
	return string(b), err
}

func (r *Res) Bytes() ([]byte, error) {
	if r.err != nil {
		return nil, r.err
	}

	b, err := ioutil.ReadAll(r.res.Body)
	if err != nil {
		return nil, err
	}
	r.res.Body.Close()
	return b, nil
}
