package main

import "fmt"

func main() {

	//Simple iteration over a range
	for i:=1; i<=5; i++ {
		fmt.Println(i)
	}

	//Iterate over collection
	numbers := [] int {1,2,3,4,5,6}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//Break and continue
	for i:=1; i<=10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number: ", i)
		if i == 5 {
			break
		}
	}

	//Nested loop
	for i:=1; i<=3; i++ {
		for j:=1; j<=3; j++ {
			fmt.Println(i, j)
		}
	}

	for i := range 10 {
		i++
		fmt.Println(i)
	}


}