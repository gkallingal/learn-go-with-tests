package selects

import (
	"net/http"
	"time"
)

func WebsiteRacer(x, y string) string {
	durationX := measureResponse(x)
	durationY := measureResponse(y)

	if durationX < durationY {
		return x
	}
	return y
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
