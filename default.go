package main

import (
	"fmt"
	"time"
)

//kondisi default pada sebuah select akan dieksekusi ketika tidak ada kondisi lain yang terpenuhi

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("     ")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
