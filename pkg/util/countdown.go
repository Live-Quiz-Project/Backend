package util

import (
	"sync"
	"time"
)

func Countdown(seconds int) <-chan int {
	timeLeft := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := seconds; i >= 0; i-- {
			timeLeft <- i
			time.Sleep(time.Second)
		}
		close(timeLeft)
	}()

	go func() {
		wg.Wait()
	}()

	return timeLeft
}
