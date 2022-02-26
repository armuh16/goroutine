package main

import "fmt"

func main() {
	c := make(chan bool, 5)
	c <- true
	// channel with value 0 of buffer calling Rendezvous Channel
	fmt.Println("this line be printed!")
}
