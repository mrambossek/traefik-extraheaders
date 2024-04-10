package traefik_extraheaders_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	extraheaders "github.com/mrambossek/traefik-extraheaders"
)

func TestExtraheaders(t *testing.T) {
	cfg := extraheaders.CreateConfig()
	cfg.ClientPortHeader = "Test-Forwarded-Clientport"
	cfg.HTTPVerHeader = "Test-Forwarded-HTTP-Ver"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := extraheaders.New(ctx, next, cfg, "demo-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	fmt.Printf("req = %v", req)
	// assertHeader(t, req, "X-Extraheaders", "test")
}

/*
func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
*/
