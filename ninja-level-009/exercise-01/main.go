package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("goroutines start:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())

	const c int = 2
	var wg sync.WaitGroup
	wg.Add(c)

	go func() { fmt.Println("anonFunc1"); wg.Done() }()
	go func() { fmt.Println("anonFunc2"); wg.Done() }()

	fmt.Println("goroutines before wait:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())

	wg.Wait()
	fmt.Println("goroutines end:", runtime.NumGoroutine())
	fmt.Println("CPUs", runtime.NumCPU())
}
