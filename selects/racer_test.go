package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	t.Run("test between two sites", func(t *testing.T) {

		slowServer := startDelayServer(10 * time.Millisecond)
		fastServer := startDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL

		got, err := WebsiteRacer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect error but got %v", err)
		}

		if got != want {
			t.Errorf("Got %q, Want %q", got, want)
		}
	})

	t.Run("test when response is greater that 10s", func(t *testing.T) {
		serverX := startDelayServer(25 * time.Millisecond)

		defer serverX.Close()

		_, err := ConfigurableRacer(serverX.URL, serverX.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected error and got none")
		}
	})

}

func startDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
