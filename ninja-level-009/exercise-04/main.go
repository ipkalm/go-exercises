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
	var mtx sync.Mutex

	counter := 0
	wg.Add(c)
	for i := 0; i < 100; i++ {
		go func() {
			mtx.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("counter:", counter)
			mtx.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("goroutines mid:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())

	wg.Wait()

	fmt.Println("counter:", counter)
	fmt.Println("goroutines end:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
}
