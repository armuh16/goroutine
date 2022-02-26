package main

import "fmt"

func main() {
	c := make(chan bool)
	c <- true

	fmt.Println("this is will never be printed!")
}
