package main

import "fmt"

func main() {
	// 1. For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	//fmt.Println(sum)

	// The init and post statements are optional.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	//fmt.Println(sum)

	// 2. While
	// At that point you can drop the semicolons: C's while is spelled for in Go.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
