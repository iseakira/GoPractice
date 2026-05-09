package main

import (
	"time"
)

func sendMessage(msg string) {
	println(msg)
}

func main() {
	message := "hi"
	go func() {
		sendMessage(message)
	}()

	message = "ho"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
}




