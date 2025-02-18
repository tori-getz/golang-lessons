package sumslice

import "sync"

func SumSliceSingleThread(arr []int) int {
	sum := 0

	for _, value := range arr {
		sum = sum + value
	}

	return sum
}

func SumSliceMultiThread(arr []int) int {
	numGoroutines := 3
	ch := make(chan int)
	partSize := len(arr) / numGoroutines

	var wg sync.WaitGroup

	sum := 0

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			start := i * partSize
			end := start + partSize
			go sumPart(arr[start:end], ch, &wg)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		sum = sum + value
	}

	return sum
}

func sumPart(arr []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0

	for _, value := range arr {
		sum = sum + value
	}

	ch <- sum
}
