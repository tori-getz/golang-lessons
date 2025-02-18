package main

import (
	"fmt"
	"goroutines/app/google"
	"sync"
	"time"
)

func main() {
	syncRun()
	asyncRun()
}

func syncRun() {
	t := time.Now()

	for i := 0; i < 9; i++ {
		google.GetMainPage(nil)
	}

	fmt.Println("Time: ", time.Since(t))
}

func asyncRun() {
	t := time.Now()

	var wg sync.WaitGroup

	statusCode := make(chan int)

	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func() {
			google.GetMainPage(statusCode)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(statusCode)
	}()

	for res := range statusCode {
		fmt.Println("Status code: ", res)
	}

	fmt.Println("Time: ", time.Since(t))
}
