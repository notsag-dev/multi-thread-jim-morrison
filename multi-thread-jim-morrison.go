package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Printf("Why did multi-thread Jim Morrison cross the road?\n")
	response := [8]string{"To", "break", "on", "through", "to", "the", "other", "side"}
	channel := make(chan string)
	for _, v := range response {
		go waitAndSend(v, channel)
	}
	for i := 0; i < 8; i++ {
		fmt.Printf("%v ", <-channel)
	}
	fmt.Println();
}

func waitAndSend(text string, output chan<- string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	output <- text
}
