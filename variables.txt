package main

import "fmt"

var middleName = "MacGuirre"
func main() {
	fmt.Println("Hello, World!")

	var name string = "John"
	fmt.Println(name)

	name1:= "John Doe"
	fmt.Println(name1)

	fmt.Println(middleName)

	middleName = "Mac"
	fmt.Println(middleName)

	test()
}

func test() {
	var name string = "John"
	fmt.Println(name)
}