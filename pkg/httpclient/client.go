package httpclient

import (
	"github.com/go-resty/resty/v2"
)

// NewRequest ...
func NewRequest() *resty.Request {
	return NewClient().R()
}

// NewClient ...
func NewClient() *resty.Client {
	c := resty.New()
	c.SetDebug(true)
	return c
}
