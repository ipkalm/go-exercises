package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("goroutines start:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Printf("\n---\n")

	var wg sync.WaitGroup
	var counter int64

	const c int = 100

	wg.Add(c)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	fmt.Println("goroutines mid:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Printf("---\n")

	wg.Wait()

	fmt.Println("goroutines end:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Printf("---\n")

	fmt.Println("counter:", counter)
}
