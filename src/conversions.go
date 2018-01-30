package main

import "fmt"

func main () {
	var value1 int
	var value2 float32

	value1 = 50

	value2 = float32(value1)

	fmt.Println(value1)
	fmt.Println(value2) 
}
