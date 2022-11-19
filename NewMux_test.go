package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	set := NewMux()
	set.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}

	got, ett := io.ReadAll(resp.Body)
	if ett != nil {
		t.Fatalf("failed to read body: %v", ett)
	}

	want := `{"status":"ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
