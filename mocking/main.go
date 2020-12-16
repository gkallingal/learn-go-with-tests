package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	startCountdown = 3
	finalWord      = "Go!"
	sleep          = "sleep"
	write          = "write"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type RealSleeper struct {
}

type CountdownOperationsSpy struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

// func (r *RealSleeper) Sleep() {
// 	time.Sleep(time.Second)
// }

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(writer io.Writer, sleep Sleeper) {
	for i := startCountdown; i > 0; i-- {
		sleep.Sleep()
		fmt.Fprintf(writer, "%v\n", i)
	}

	sleep.Sleep()
	fmt.Fprintf(writer, "%v\n", finalWord)

}

func main() {
	sleeptime := 5 * time.Second
	sleeper := &ConfigurableSleeper{sleeptime, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
