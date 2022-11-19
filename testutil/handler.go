package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJson(t *testing.T, want, got []byte) {
	t.Helper()

	var expectedJson, actualJson any
	if err := json.Unmarshal(want, &expectedJson); err != nil {
		t.Fatalf("failed to unmarshal expected json: %v", err)
	}
	if err := json.Unmarshal(got, &actualJson); err != nil {
		t.Fatalf("failed to unmarshal actual json: %v", err)
	}

	if diff := cmp.Diff(expectedJson, actualJson); diff != "" {
		t.Errorf("json is not equal (-want +got): %s", diff)
	}
}

func AssertJsonString(t *testing.T, want, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close() })
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}

	if got.StatusCode != status {
		t.Errorf("want status %d, but got %d, body:%d", status, got.StatusCode, gb)
	}

	if len(gb) == 0 && len(body) == 0 {
		// 期待としても実態としてもレスポンスボディがないため
		// AssertJsonの処理は不要
		return
	}
	AssertJson(t, body, gb)
}
