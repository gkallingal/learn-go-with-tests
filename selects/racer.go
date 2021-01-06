package selects

import (
	"net/http"
)

func WebsiteRacer(x, y string) string {
	select {
	case <-ping(x):
		return x
	case <-ping(y):
		return y
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
