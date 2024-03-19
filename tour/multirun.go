package main

import (
	"fmt"
	"time"
)

// Goroutines:
// A goroutine is a thread managed by runtime.
// Goroutines run in the same address space which brings the need of synchronization.
// go f(x, y, z), sharing same address space
// thread 1							thread 2
// evaluation of f, x, y, and z     execution of f
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Channels: a conduit that data flows instead of water
// representation: ch <- v, v := <-ch
// initialization: ch := make(chan int)
// Sends and receives block until the other side is ready to receive/send.
// One goroutine sends a value to the channel. The send op on the channel waits for another goroutine receiving it.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// Buffered channels
// buf_ch := make(chan int, 100), providing buffer length
// Sends block only when the buffer is full. Receives block only when the buffer is empty.

// Range and close:
// v, ok := <-ch, ok indicates if any value is still in the channel
// for i := range ch, the loop goes on until the channel is closed
// only the sender can close the channel. Sends on a closed channel panic
func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Select
func fib_select(c, quit chan int) {
	x, y := 0, 1
	for {
		// Select blocks until one of its case can run. Random selection if multiple are ready
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	//// When creating a goroutine, it runs concurrently with the rest of the program.
	//go say("world")
	//// Main() doesn't care if goroutines finish and returns anyway. When main() returns, the program exits despite
	//// running goroutines
	//say("hello")
	//
	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//// the order of sends and receives is not guaranteed unless specified
	//x, y := <-c, <-c // receive from c
	//fmt.Println(x, y, x+y)

	//buf_ch := make(chan int, 2)
	//buf_ch <- 1
	//buf_ch <- 8
	//buf_ch <- 5 // buffer overflow
	//fmt.Println(<-buf_ch)
	//fmt.Println(<-buf_ch)
	//fmt.Println(<-buf_ch)

	//c := make(chan int, 10)
	//go fib(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib_select(c, quit)

	// Default selection
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		// in default case receives block
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
