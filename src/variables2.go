package main

import "fmt"

func main() {

	// a
	var str = "test"
	var integer = 10
	var boolean = true

	fmt.Printf("str = %s\n", str)
	fmt.Printf("integer = %d\n", integer)
	fmt.Println(boolean)

	fmt.Printf("str = %s\n", str)
	fmt.Printf("integer = %d\n", integer)
	fmt.Println(boolean)

	// b
	str1 := "test"
	integer1 := 10
	boolean1 := true

	fmt.Printf("str1 = %s\n", str1)
	fmt.Printf("integer1 = %d\n", integer1)
	fmt.Println(boolean1)
}
