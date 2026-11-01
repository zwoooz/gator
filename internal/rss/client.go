// Package rss
// WARN: above is linting supresson so it doesn't mark everything in package rss
// write and actual description here later and remove this warning
package rss

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
