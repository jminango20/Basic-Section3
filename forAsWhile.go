package main

import "fmt"
import "time"
import "math/rand"

func main() {

	/*
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	//While
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	*/

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a random number between 1 to 100
	target := r.Intn(100) + 1

	// Welcome message
	fmt.Println("Guess the number between 1 to 100")
	fmt.Print("Your guess: ")

	var guess int

	for {
		fmt.Println("Choose a number")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("You guessed it!")
			break
		} else if guess > target {
			fmt.Println("Too high")
		} else {
			fmt.Println("Too low")
		}
	}

}