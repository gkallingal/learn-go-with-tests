package selects

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func WebsiteRacer(x, y string) (string, error) {
	return ConfigurableRacer(x, y, tenSecondTimeout)

}

func ConfigurableRacer(x, y string, timeout time.Duration) (string, error) {
	select {
	case <-ping(x):
		return x, nil
	case <-ping(y):
		return y, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timeout waiting for %s and %s to respond", x, y)
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
