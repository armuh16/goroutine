package main

import "fmt"

func main() {
	anon := []int{1, 2, 3, 4, 5}
	for _, v := range anon {
		go func(i int) {
			//can do anything else
			fmt.Println(i)
		}(v)
	}
}
