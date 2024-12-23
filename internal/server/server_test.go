package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	srv := NewServer()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.handleHealth()(rec, req)

	resp := rec.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.StatusCode)
	}

	if string(body) != "OK" {
		t.Errorf("expected body 'OK'; got %v", string(body))
	}
}

func TestServer_Start(t *testing.T) {
	srv := NewServer()
	go func() {
		err := srv.Start(":0")
		if err != http.ErrServerClosed {
			t.Errorf("expected ErrServerClosed; got %v", err)
		}
	}()
}
