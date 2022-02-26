package main

import (
	"fmt"
	"time"
)

// goroutine lightweight thread yang dikelola oleh runtime go
// go (x, y, z)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}
