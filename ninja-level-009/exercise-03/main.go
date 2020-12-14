package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("goroutines start:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())

	var wg sync.WaitGroup
	const c int = 100

	counter := 0
	wg.Add(c)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	fmt.Println("goroutines mid:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())

	wg.Wait()

	fmt.Println("goroutines end:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
}
