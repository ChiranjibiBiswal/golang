package main

import (
	"fmt"
	"time"
)

// printData receives one value from the channel.
func printData(name <-chan string) {
	fmt.Println("Printing channel data:", <-name)
}

// printNumberData receives values until the channel is closed.
func printNumberData(nums <-chan int) {
	for num := range nums {
		fmt.Println("Printing number data:", num)
	}
}

// sum sends the sum to the result channel.
func sum(result chan<- int, num1, num2 int) {
	result <- num1 + num2
}

// taskStatus sends a completion signal when the task finishes.
func taskStatus(task chan<- bool) {
	defer func() {
		task <- true
	}()

	fmt.Println("Processing...")
	time.Sleep(time.Millisecond * 100)
}

// emailSending processes emails and reports completion.
func emailSending(emails <-chan string, done chan<- bool) {
	defer func() {
		done <- true
	}()

	for email := range emails {
		fmt.Println("Email:", email)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// Unbuffered channel: the send waits until printData receives.
	messageChan := make(chan string)
	go printData(messageChan)
	messageChan <- "Aman"

	// Closed channel: printNumberData receives values until the channel closes.
	numberChan := make(chan int)
	go printNumberData(numberChan)
	for i := 1; i <= 5; i++ {
		numberChan <- i
	}
	close(numberChan)

	// Send a calculated result through a channel.
	resultChan := make(chan int)
	go sum(resultChan, 10, 5)
	fmt.Println("Printing result:", <-resultChan)

	// Unidirectional channel: wait for a task completion signal.
	taskChan := make(chan bool)
	go taskStatus(taskChan)
	fmt.Println("Task complete:", <-taskChan)

	// Buffered channel: values can be queued before receiving.
	emailChan := make(chan string, 5)
	doneChan := make(chan bool)
	go emailSending(emailChan, doneChan)

	for i := 1; i <= 5; i++ {
		emailChan <- fmt.Sprintf("email%d@gmail.com", i)
	}
	close(emailChan)
	fmt.Println("Done sending:", <-doneChan)

	// Select: receive whichever channel is ready first.
	numberValueChan := make(chan int)
	nameValueChan := make(chan string)

	go func() {
		numberValueChan <- 1
	}()

	go func() {
		nameValueChan <- "Goku"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case numberValue := <-numberValueChan:
			fmt.Println("Channel 1 value:", numberValue)
		case nameValue := <-nameValueChan:
			fmt.Println("Channel 2 value:", nameValue)
		}
	}
}
