package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Test countdown from 3 to Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &CountdownOperationsSpy{}
		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("Got %s, Want %s", got, want)
		}
	})
	t.Run("test order of sleep and write", func(t *testing.T) {
		sleeper := &CountdownOperationsSpy{}
		Countdown(sleeper, sleeper)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, sleeper.Calls) {
			t.Errorf("Got %v, want %v", sleeper.Calls, want)
		}
	})
	t.Run("test configurable sleeper", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("%v vs %v", sleepTime, spyTime.durationSlept)
		}
	})
}
