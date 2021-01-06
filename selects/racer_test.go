package selects

import "testing"

func TestWebsiteRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.google.com"

	want := fastURL

	got := WebsiteRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("Got %q, Want %q", got, want)
	}
}
