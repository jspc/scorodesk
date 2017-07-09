package main

import (
	"io"
	"net/http"
	"testing"
)

type MockHTTP struct{}

func (c MockHTTP) Post(url, cType string, i io.Reader) (r *http.Response, err error) { return }

func TestRequest(t *testing.T) {
	for _, test := range []struct {
		name        string
		module      string
		action      string
		data        interface{}
		expectError bool
		client      httpClient
	}{
		{"some request", "contact", "add", []string{}, false, MockHTTP{}},
	} {
		t.Run(test.name, func(t *testing.T) {
			s := Scoro{"some-key", "jspc", test.client}

			_, err := s.request(test.module, test.action, test.data)

			if err == nil {
				if test.expectError {
					t.Errorf("expected error, received none")
				}
			} else {
				if !test.expectError {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestURL(t *testing.T) {
	for _, test := range []struct {
		name    string
		company string
		module  string
		action  string
		expect  string
	}{
		{"simple url", "company", "module", "action", "https://company.scoro.ee/api/module/action"},
	} {
		t.Run(test.name, func(t *testing.T) {
			s := Scoro{"some-key", test.company, MockHTTP{}}
			u := s.url(test.module, test.action)

			if test.expect != u {
				t.Errorf("expected %q, received %q", test.expect, u)
			}
		})
	}
}
