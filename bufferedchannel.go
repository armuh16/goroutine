package main

import "fmt"

// channel dapat memilih buffer.
// - pengiriman data terkunci ketika channel penuh
// - penerimaan data terkunci ketika channel kosong

func main() {
	c := make(chan bool, 5)
	c <- true
	// channel with value 0 of buffer calling Rendezvous Channel
	fmt.Println("this line be printed!")
}
