package main

import "fmt"

// secara defalut, pengirim(send) dan penerima(receive) akan terkunci/stagnan hingga sisi yang lain siap untuk menerima atau mengirim data

func main() {
	c := make(chan bool)
	c <- true

	fmt.Println("this is will never be printed!")
}
