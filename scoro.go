package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type httpClient interface {
	Post(string, string, io.Reader) (*http.Response, error)
}

type Scoro struct {
	APIKey      string
	CompanyName string

	client httpClient
}

func NewScoro(name, apiKey string) (s Scoro) {
	s.APIKey = apiKey
	s.CompanyName = name

	return
}

func (s Scoro) request(module, action string, data interface{}) (resp *http.Response, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)
	return http.Post(s.url(module, action), "application/json", buf)
}

func (s Scoro) url(module, action string) string {
	return fmt.Sprintf("https://%s.scoro.ee/api/%s/%s",
		s.CompanyName,
		module,
		action,
	)
}
