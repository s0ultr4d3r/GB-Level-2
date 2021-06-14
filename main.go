package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var sum = 0

func process(n string) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 10000; i++ {
			mu.Lock()
			sum = sum + 1
			mu.Unlock()
		}

		fmt.Println("From "+n+":", sum)
	}()
}

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	processes := []string{"A", "B", "C", "D", "E"}
	for _, p := range processes {
		process(p)
	}

	wg.Wait()
	fmt.Println("Final Sum:", sum)
}
