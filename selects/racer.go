package selects

import (
	"net/http"
	"time"
)

func WebsiteRacer(x, y string) string {
	startX := time.Now()
	http.Get(x)
	durationX := time.Since(startX)

	startY := time.Now()
	http.Get(y)
	durationY := time.Since(startY)

	if durationX < durationY {
		return x
	}
	return y
}
