package client

import (
	"net/http"
	"time"
)

// TODO once
func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
}
