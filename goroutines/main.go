package main

import (
	"fmt"
	"goroutines/app/google"
	"goroutines/app/sumslice"
	"sync"
	"time"
)

func main() {
	googleSingleThread()
	googleMultiThread()
	sumSliceSingleThread()
	sumSliceMultiThread()
}

func googleSingleThread() {
	println("Google.com request (Single thread)")
	t := time.Now()

	for i := 0; i < 9; i++ {
		google.GetMainPage(nil)
	}

	fmt.Println("Time: ", time.Since(t))
	println("")
}

func googleMultiThread() {
	println("Google.com request (Multi thread)")
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
	println("")
}

var sliceForTest = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func sumSliceSingleThread() {
	println("Sum slice (Single thread)")
	t := time.Now()

	sum := sumslice.SumSliceSingleThread(sliceForTest)

	fmt.Println("Sum: ", sum)
	fmt.Println("Time: ", time.Since(t))
	println("")
}

func sumSliceMultiThread() {
	println("Sum slice (Multi thread)")
	t := time.Now()

	sum := sumslice.SumSliceMultiThread(sliceForTest)

	fmt.Println("Sum: ", sum)
	fmt.Println("Time: ", time.Since(t))
	println("")
}
