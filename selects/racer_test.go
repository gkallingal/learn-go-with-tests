package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {

	slowServer := startDelayServer(10 * time.Millisecond)
	fastServer := startDelayServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL

	got := WebsiteRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("Got %q, Want %q", got, want)
	}

}

func startDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
