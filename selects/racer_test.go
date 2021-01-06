package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got := WebsiteRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("Got %q, Want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}
