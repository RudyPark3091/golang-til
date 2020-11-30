package main

import (
	"fmt"
	"sync"
	"time"
)

// function waits sec second and print one line
func waitAndPrint(sec time.Duration) {
	time.Sleep(sec * time.Second)
	fmt.Printf("waited %d second\n", sec)
}

// function for goroutine
func waitAndPrintGoroutine(sec time.Duration) {
	time.Sleep(sec * time.Second)
	fmt.Printf("goroutine watied %d second\n", sec)
}

func waitAndPrintWG(sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(sec * time.Second)
	fmt.Printf("WG waited %d second\n", sec)
}

func WaitingFunctionExample() {
	iteration := 5

	// function calls not using goroutine
	// one line per second
	for i := 0; i < iteration; i++ {
		waitAndPrint(1)
	}

	// function calls using goroutine
	// every line printed at once
	for i := 0; i < iteration; i++ {
		go waitAndPrintGoroutine(1)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Waiting finished")
}

func WaitingUsingWaitGroup() {
	iteration := 5
	// using WaitGroup Struct defined in sync package
	wg := new(sync.WaitGroup)

	for i := 0; i < iteration; i++ {
		// wg.Add should be called Every goroutine
		wg.Add(1)
		// function synchronized with wg must call wg.Done()
		// if not: main process emits deadlock error
		go waitAndPrintWG(1, wg)
	}

	// wait until all added goroutines return
	wg.Wait()
}

func main() {
	WaitingFunctionExample()
	WaitingUsingWaitGroup()
}
