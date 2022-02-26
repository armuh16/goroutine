package main

import "fmt"

// function yang didefinisikan tanpa nama fungsi yang mengacu ke fungsi tersebut

func main() {
	anon := []int{1, 2, 3, 4, 5}
	for _, v := range anon {
		go func(i int) {
			//can do anything else
			fmt.Println(i)
		}(v)
	}
}
