package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

// If we want to send/receive data from one go routine to another, we use channels
// Channels are used to do communication between different go routines

// Deadlock
// Channels are blocking till other side is not ready to receive

// func main() {
// 	messageChan := make(chan string) // Channel creation

// 	messageChan <- "ping" // Sending data

// 	msg := <- messageChan // Receiving data

// 	fmt.Println(msg)
// }

// Sending data
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// 	// fmt.Println("Processing number", <-numChan)
// }

// func main() {
// 	numChan := make(chan int)

// 	go processNum(numChan)

// 	for { // Infinite loop
// 		numChan <- rand.Intn(100)
// 	}

// 	// numChan <- 5
// 	// time.Sleep(time.Second * 2)
// }

// Receiving data
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func main() {
// 	result := make(chan int)

// 	go sum(result, 4, 5)

// 	res := <- result // Blocking (Partially true)

// 	fmt.Println(res)
// }

// Go routine synchronizer - Channel vs WaitGroup
// func task(done chan bool) {
// 	defer func() { done <- true }() // Executes after function execution ends
// 	fmt.Println("Processing...")
// }

// func main() {
// 	done := make(chan bool)

// 	go task(done)

// 	<- done // Block
// }

// Buffered channel + Type safety
// func emailSender(
// 	emailChan <-chan string, // Receive only channel
// 	done chan<- bool, // Send only channel
// ) {
// 	defer func() { done <- true }()

// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	emailChan := make(chan string, 100)
// 	done := make(chan bool)

// 	go emailSender(emailChan, done)

// 	for i := 0; i < 10; i++ {
// 		emailChan <- fmt.Sprintf("example%d@gmail.com", i)
// 	}

// 	fmt.Println("Done email sending")

// 	close(emailChan) // Important
// 	<- done

// 	// emailChan <- "example1@gmail.com"
// 	// emailChan <- "example2@gmail.com"

// 	// fmt.Println(<- emailChan)
// 	// fmt.Println(<- emailChan)
// }

// Listen to multiple channels
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <- chan1:
			fmt.Println("Received data from chan1", chan1Val)
		case chan2Val := <- chan2:
			fmt.Println("Received data from chan2", chan2Val)
		}
	}
}
