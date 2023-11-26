package main

import (
	"fmt"
	"time"
)

func writePing(word chan string) {
	for i := 0; ; i++ {
		word <- "ping"
	}
}

func writePong(word chan string) {
	for i := 0; ; i++ {
		word <- "pong"
	}
}

func WriteMsgs(word chan string) {
	for {
		select {
		case msg := <-word:
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	var word chan string = make(chan string)

	go writePing(word)
	go WriteMsgs(word)
	go writePong(word)

	var input string
	fmt.Scanln(&input)
}
