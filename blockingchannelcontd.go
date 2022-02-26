package main

import "fmt"

//menjalankan goroutine untuk mengkonsumsi data dari channel c

func main() {
	c := make(chan bool)
	go func() {
		<-c
	}()
	c <- true

	fmt.Println("this line is printed!")
}
