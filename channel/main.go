package main

import (
	"fmt"
	"time"
)

func BasicChannelExample() {
	// defining channel
	// needs type assertion
	// usually channel is for communication between goroutines
	ch := make(chan int)

	// closing channel
	// can't send data anymore still can receive data from channel
	defer close(ch)

	/* main routine emits deadlock error if used with no goroutines
	 *
	 * func() {
	 * 	ch <- 1
	 * }()
	 *
	 */

	// send data into channel
	go func() {
		ch <- 1
	}()

	// receive data from channel
	number := <-ch
	fmt.Println(number)
}

func ChannelWithGoroutine() {
	end := make(chan bool)
	defer close(end)

	// anonymous function goroutine
	// can wait goroutine without using sync.WaitGroup with channel
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%dth goroutine iteration\n", i+1)
			time.Sleep(200 * time.Millisecond)
		}
		// to finish, send something on channel
		end <- true
	}()

	// waiting goroutine above to finish with this expression
	// <-"channel name"
	<-end
}

func BufferedChannel() {
	// unbuffered channel
	// ch := make(chan int)
	// emits error: no goroutines
	// ch <- 1
	// unbuffered channel needs receiver goroutine
	// but there's no receiver if there's no goroutine
	// eventually it emits error

	// buffered channel
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	iteration := 3
	for i := 0; i < iteration; i++ {
		fmt.Printf("received through buffered channel %d\n", <-ch)
		time.Sleep(time.Millisecond * 200)
	}
}

// receive-only channel
// <-chan (type)
func receiverChan(ch <-chan string) {
	fmt.Printf("received through channel %s\n", <-ch)

	// error when sending data into channel
	// ch <- "Hello"
}

// send-only channel
// chan<- (type)
func sendChan(ch chan<- string, message string) {
	ch <- message

	// error when receiving data from channel
	// received := <- ch
}

func ChannelDirection() {
	// string type channel
	msg := make(chan string, 1)
	sendChan(msg, "Hello")
	received := <-msg
	fmt.Printf("received message: %s\n", received)
}

// Check if there's data left in channel and flushing it
func FlushChannelBuffer(ch chan int) {
	// <-ch returns bool whether reception is successful
	// if channel is closed and no data left in channel, it returns false
	for {
		if d, left := <-ch; left {
			fmt.Printf("remaining data : %d\n", d)
		} else {
			fmt.Println("Channel Closed\nNo more data left in channel")
			break
		}
	}

	// same expression as above
	// for d := range ch {
	// 	fmt.Println(d)
	// }
}

func Flush_Test() {
	itr := 10
	ch := make(chan int, itr)

	for i := 0; i < itr; i++ {
		ch <- i + 1
	}
	close(ch)

	FlushChannelBuffer(ch)
}

func chan1(ch chan bool) {
	time.Sleep(200 * time.Millisecond)
	ch <- true
}

func chan2(ch chan bool) {
	time.Sleep(300 * time.Millisecond)
	ch <- true
}

func SelectExp() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go chan1(ch1)
	go chan2(ch2)

	// using LABEL in go
	// no line should be in between LABEL and loop
LOOP:
	for {
		select {
		case <-ch1:
			fmt.Println("chan1 exit")

		case <-ch2:
			fmt.Println("chan2 exit")
			// break LABEL
			// exits loop and execute next line of the loop block
			break LOOP
		}
	}
}

func main() {
	BasicChannelExample()
	ChannelWithGoroutine()
	BufferedChannel()
	ChannelDirection()
	Flush_Test()
	SelectExp()
}
